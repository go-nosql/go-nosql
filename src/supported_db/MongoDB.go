package supported_db

import (
	"db/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"strconv"
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

// First - Read first record from mongoDB 
func (this MongoDb) First() entity.Map {
        var record entity.Map
        this.Conn.Find(nil).One(&record)
        return record
}

// Last - Read last record from mongoDB 
func (this MongoDb) Last() entity.Map {
        var record entity.Map
        count, _ := this.Conn.Count()
        this.Conn.Find(nil).Skip(count-1).One(&record)
        return record
}

// Count - Read number of records from mongoDB
func (this MongoDb) Count() int {
        count, _ := this.Conn.Count()
        return count
}

// Limit - Read limited number of records from mongoDB.
func (this MongoDb) Limit(limit int) []entity.Map {
        var records []entity.Map
	if limit <= 0 {
		return records
	}
        this.Conn.Find(nil).Limit(limit).All(&records)
        return records
}

// Where - Get the records based on query string.
func (this MongoDb) Where(query string) []entity.Map {
        var records []entity.Map
	all := strings.Fields(query)
	value := all[2]
	var val interface{}
	if string(value[0]) == "'" && string(value[len(value)-1]) == "'" {
	 val = value[1:(len(value)-1)]
	} else {
	 val, _ = strconv.ParseFloat(value,64)
	}
	switch all[1] {
		case "<" : this.Conn.Find(bson.M{all[0]:bson.M{"$lt":val}}).All(&records)
		case ">" : this.Conn.Find(bson.M{all[0]:bson.M{"$gt":val}}).All(&records)
		case "=", "==" : this.Conn.Find(bson.M{all[0]:val}).All(&records)
		case "!=" : this.Conn.Find(bson.M{all[0]:bson.M{"$ne":val}}).All(&records)
		case "<=" : this.Conn.Find(bson.M{all[0]:bson.M{"$lte":val}}).All(&records)
		case ">=" : this.Conn.Find(bson.M{all[0]:bson.M{"$gte":val}}).All(&records)
	}
        return records
}

// FindById - Read records by id from mongoDB.
func (this MongoDb) FindById(id string) entity.Map {
        var record entity.Map
        this.Conn.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&record)
        return record
}
