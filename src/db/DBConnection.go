package db

import (
	"configparser-master"
	"couch-go-master"
	"fmt"
	"gopkg.in/mgo.v2"
	"strings"
	"supported_db"
)

func GetConnection(path string) (supported_db.CouchDb, supported_db.MongoDb) {
	var collection *mgo.Collection
	conf, err := configparser.Read(path)
	if err != nil {
		panic(err)
	}
	config, _ := conf.Section("database")
	switch strings.ToUpper(config.ValueOf("Database")) {
	case "COUCH":
		db, _ := couch.NewDatabase(config.ValueOf("IPAddress"), config.ValueOf("Port"), config.ValueOf("DbName"))
		return supported_db.CouchDb{db}, supported_db.MongoDb{}
	case "MONGO":
		mongoSession, err := mgo.Dial(config.ValueOf("IPAddress") + ":" + config.ValueOf("Port"))
		if err != nil {
			panic(err)
		} else {
			db := mongoSession.DB(config.ValueOf("DbName"))
			collection = db.C(config.ValueOf("CollectionName"))
			fmt.Println(collection)
		}
		return supported_db.CouchDb{}, supported_db.MongoDb{collection}
	}
	return supported_db.CouchDb{}, supported_db.MongoDb{}
}
