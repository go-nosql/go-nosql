package supported_db

import (
	"db/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDB - Struct for mongo database.
type MongoDb struct {
	Conn *mgo.Collection
}

// Read - Read all records from mongoDB.
func (this MongoDb) Read() []entity.Map {
	var records []entity.Map
	this.Conn.Find(nil).All(&records)
	return records
}

// Save - Save generic record in mongoDB.
func (this MongoDb) Save(record interface{}) bool {
	err := this.Conn.Insert(record)
	if err == nil {
		return true
	} else {
		return false
	}
}

// Delete - Delete generic record in mongoDB.
func (this MongoDb) Delete(record map[string]interface{}) bool {
	err := this.Conn.Remove(bson.M{"_id": record["_id"]})
	if err == nil {
		return true
	} else {
		return false
	}
}

// Update - Update record in mongoDB.
func (this MongoDb) Update(record map[string]interface{}) bool {
	err := this.Conn.UpdateId(record["_id"], record)
	if err == nil {
		return true
	} else {
		return false
	}
}
