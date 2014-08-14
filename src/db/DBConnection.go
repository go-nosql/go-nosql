package db

import (
	"configparser-master"
	"couch-go-master"
	"gopkg.in/mgo.v2"
	"strings"
	"supported_db"
	"abstract"
)

func GetConnection(path string) abstract.Database {
	var collection *mgo.Collection
	conf, err := configparser.Read(path)
	if err != nil {
		panic(err)
	}
	driver_config, _ := conf.Section("database")
	db_config, _ := conf.Section(driver_config.ValueOf("name"))
	switch strings.ToUpper(driver_config.ValueOf("name")) {
	case "COUCH":
		var db couch.Database
		if db_config.ValueOf("user") == ""{
			db, _ = couch.NewDatabase(db_config.ValueOf("ipaddress"), db_config.ValueOf("port"), db_config.ValueOf("dbname"))
		} else {
			db, _ = couch.NewDatabaseByURL("http://" + db_config.ValueOf("user") + ":" + db_config.ValueOf("password") + "@" + db_config.ValueOf("ipaddress") + ":" + db_config.ValueOf("port") + "/" + db_config.ValueOf("dbname"))
		}
		return abstract.Database(supported_db.CouchDb{db})
	case "MONGO":
		var mongoSession *mgo.Session
		if db_config.ValueOf("user") == "" {
			mongoSession, err = mgo.Dial(db_config.ValueOf("ipaddress") + ":" + db_config.ValueOf("port"))
		} else {
			mongoSession, err = mgo.Dial(db_config.ValueOf("user") + ":" + db_config.ValueOf("password") + "@" + db_config.ValueOf("ipaddress") + ":" + db_config.ValueOf("port") + "/" + db_config.ValueOf("dbname"))
		}
		if err != nil {
			panic(err)
		} else {
			db := mongoSession.DB(db_config.ValueOf("dbname"))
			collection = db.C(db_config.ValueOf("collectionname"))
		}
		return abstract.Database(supported_db.MongoDb{collection})
	}
	return nil
}
