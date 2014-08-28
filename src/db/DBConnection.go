package db

import (
	"abstract"
	"configparser-master"
	"couch-go-master"
	"db/entity"
	"gopkg.in/mgo.v2"
	"reflect"
	"strings"
	"supported_db"
)

// GetConnection - Return DB connection object based on config file.
func GetConnection(path string) abstract.Database {
	var collection *mgo.Collection
	conf, _ := configparser.Read(path)
	checkError(conf, "Invalid configuration file")
	driver_config, _ := conf.Section("nosql.db")

	var db_config *configparser.Section
	if strings.Contains(strings.ToLower(driver_config.ValueOf("name")), "mongo") {
		db_config, _ = conf.Section("mongo")
		driver_config.SetValueFor("name", "mongo")
	} else if strings.Contains(strings.ToLower(driver_config.ValueOf("name")), "couch") {
		db_config, _ = conf.Section("couch")
		driver_config.SetValueFor("name", "couch")
	}
	switch strings.ToUpper(driver_config.ValueOf("name")) {
	case "COUCH":
		var db couch.Database
		if db_config.ValueOf("user") == "" {
			db, _ = couch.NewDatabase(db_config.ValueOf("ipaddress"), db_config.ValueOf("port"), db_config.ValueOf("dbname"))
		} else {
			db, _ = couch.NewDatabaseByURL("http://" + db_config.ValueOf("user") + ":" + db_config.ValueOf("password") + "@" + db_config.ValueOf("ipaddress") + ":" + db_config.ValueOf("port") + "/" + db_config.ValueOf("dbname"))
		}
		return abstract.Database(supported_db.CouchDb{db})
	case "MONGO":
		var mongoSession *mgo.Session
		var err error
		if db_config.ValueOf("user") == "" {
			mongoSession, err = mgo.Dial(db_config.ValueOf("ipaddress") + ":" + db_config.ValueOf("port"))
		} else {
			mongoSession, err = mgo.Dial(db_config.ValueOf("user") + ":" + db_config.ValueOf("password") + "@" + db_config.ValueOf("ipaddress") + ":" + db_config.ValueOf("port") + "/" + db_config.ValueOf("dbname"))
		}
		checkError(mongoSession, err)
		db := mongoSession.DB(db_config.ValueOf("dbname"))
		collection = db.C(db_config.ValueOf("collectionname"))
		return abstract.Database(supported_db.MongoDb{collection})
	default:
		panic("Supports only Couch or MongoDb")
	}
	return nil
}

//NewObject - Create and return a new object of type Map.
func NewObject() entity.Map {
	return make(entity.Map)
}

// NewJson - Create and return a new object of type Json.
// Ex. jsn := NewJson(`{"name":"abc","age":24}`)
// or
// jsn := NewJson()
// jsn = `{"name":"abc","age":24}`
func NewJson(s ...string) entity.Json {
	if len(s) == 0 {
		return entity.Json("")
	}
        return entity.Json(s[0])
}

// checkError - Private function to check errors.
func checkError(object interface{}, errMsg interface{}) {
	if object == nil || reflect.ValueOf(object).IsNil() {
		panic(errMsg)
	}
}
