package supported_db

import (
	"db/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"strconv"
	"reflect"
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
	if record == nil {
		return false
	}
        if reflect.TypeOf(record).String() == "string" {
                record = entity.Json(record.(string)).ToObject()
        }
	err := this.Conn.Insert(record)
	if err == nil {
		return true
	}
	return false
}

// Delete - Delete generic record in mongoDB.
func (this MongoDb) Delete(record interface{}) bool {
	if record == nil {
		return false
	}
	var err error
        if reflect.TypeOf(record).String() == "string" {
                record = entity.Json(record.(string)).ToObject()
		err = this.Conn.Remove(bson.M{"_id": bson.ObjectIdHex(record.(entity.Map)["_id"].(string))})
        } else {
		err = this.Conn.Remove(bson.M{"_id": record.(entity.Map)["_id"]})
	}
	if err == nil {
		return true
	}
	return false
}

// Update - Update record in mongoDB.
func (this MongoDb) Update(record interface{}) bool {
	if record == nil {
		return false
	}
        var err error
        if reflect.TypeOf(record).String() == "string" {
		record = entity.Json(record.(string)).ToObject()
		record.(entity.Map)["_id"] = bson.ObjectIdHex(record.(entity.Map)["_id"].(string))
        }
	err = this.Conn.UpdateId(record.(entity.Map)["_id"], record)
	if err == nil {
		return true
	}
	return false
}

// First - Read first record from mongoDB 
func (this MongoDb) First() entity.Map {
        var record entity.Map
        _ = this.Conn.Find(nil).One(&record)
        return record
}

// Last - Read last record from mongoDB 
func (this MongoDb) Last() entity.Map {
        var record entity.Map
        count, _ := this.Conn.Count()
        _ = this.Conn.Find(nil).Skip(count-1).One(&record)
        return record
}

// Count - Read number of records from mongoDB
func (this MongoDb) Count() int {
        count, err := this.Conn.Count()
	if err != nil {
		return -1
	}
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
	if len(all) < 3 {
		return records
	}
	value := strings.Join(all[2:len(all)]," ")
	var val interface{}
	var err error
	if string(value[0]) == "'" && string(value[len(value)-1]) == "'" {
	 val = value[1:(len(value)-1)]
	} else if val, err = strconv.ParseFloat(value, 64); err!=nil {
         	val, err = strconv.ParseBool(value)
        }
	switch all[1] {
		case "<" : this.Conn.Find(bson.M{all[0]:bson.M{"$lt":val}}).All(&records)
		case ">" : this.Conn.Find(bson.M{all[0]:bson.M{"$gt":val}}).All(&records)
		case "=", "==" :
				if all[0] == "_id" {
					recById := this.FindById(val.(string))
					if recById != nil {
						records = make([]entity.Map,1)
						records[0] = recById
					}
				} else {
					this.Conn.Find(bson.M{all[0]:val}).All(&records)
				}
		case "!=" : this.Conn.Find(bson.M{all[0]:bson.M{"$ne":val}}).All(&records)
		case "<=" : this.Conn.Find(bson.M{all[0]:bson.M{"$lte":val}}).All(&records)
		case ">=" : this.Conn.Find(bson.M{all[0]:bson.M{"$gte":val}}).All(&records)
	}
        return records
}

// FindById - Read record by id from mongoDB.
func (this MongoDb) FindById(id string) entity.Map {
        var record entity.Map
	defer func() {
		recover()
	}()
        this.Conn.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&record)
        return record
}
