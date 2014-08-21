package supported_db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"db/entity"
)

type MongoDb struct {
	Conn *mgo.Collection
}

func (this MongoDb) Read() []entity.Map {
	var records []entity.Map
	this.Conn.Find(nil).All(&records)
	return records
}

func (this MongoDb) Save(record interface{}) bool {
	err := this.Conn.Insert(record)
	if err == nil {
		return true
	} else {
		return false
	}
}

func (this MongoDb) Delete(record map[string]interface{}) bool {
	err := this.Conn.Remove(bson.M{"_id": record["_id"]})
	if err == nil {
		return true
	} else {
		return false
	}
}

func (this MongoDb) Update(record map[string]interface{}) bool {
	err := this.Conn.UpdateId(record["_id"], record)
	if err == nil {
		return true
	} else {
		return false
	}
}

