package db

import (
	"configparser-master"
	"couch-go-master"
	"gopkg.in/mgo.v2"
	"strings"
	"supported_db"
	"abstract"
	"os"
	"x2j-master"
)

func GetConnection(path string) (abstract.Database,map[string]interface{}) {

	var collection *mgo.Collection
	conf, err := configparser.Read(path)
	if err != nil {
		panic(err)
	}
	driver_config, _ := conf.Section("nosql.db")


	xmlFile, _ := os.Open(driver_config.ValueOf("schemalocation"))
        stat, _ := xmlFile.Stat()
        bs := make([]byte, stat.Size())
        _,_ = xmlFile.Read(bs)
        str := string(bs)
        var m map[string]interface{}
        m,_ = x2j.DocToMap(str, false)
        field := m["table"].(map[string]interface{})
        arr := field["field"].([]interface{})
        schema := make(map[string]interface{}, len(arr))
        for _, v := range arr {
        	t := v.(map[string]interface{})
        	schema[t["-name"].(string)] = t["-type"]
        }
        


	var db_config *configparser.Section
	//to make name of the database in config file case-insensitive
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
		if db_config.ValueOf("user") == ""{
			db, _ = couch.NewDatabase(db_config.ValueOf("ipaddress"), db_config.ValueOf("port"), db_config.ValueOf("dbname"))
		} else {
			db, _ = couch.NewDatabaseByURL("http://" + db_config.ValueOf("user") + ":" + db_config.ValueOf("password") + "@" + db_config.ValueOf("ipaddress") + ":" + db_config.ValueOf("port") + "/" + db_config.ValueOf("dbname"))
		}
		return abstract.Database(supported_db.CouchDb{db}),schema
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
		return abstract.Database(supported_db.MongoDb{collection}),schema
	}
	return nil,nil
}
