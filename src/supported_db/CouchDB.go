package supported_db

import (
	"couch-go-master"
	"fmt"
	"db/entity"
)

type CouchDb struct {
	Conn couch.Database
}

func (this CouchDb) Save(record interface{}) bool {
	id, rev, err := this.Conn.Insert(record)
	if err == nil && id != "" && rev != "" {
		return true
	} else {
		return false
	}
}

func (this CouchDb) Read() []entity.Map {
	ids, err := this.Conn.QueryIds("_all_docs", nil)
	records := make([]entity.Map, len(ids))
	if err != nil {
		panic(err)
	} else {
		for i := 0; i < len(ids); i++ {
			_, err = this.Conn.Retrieve(ids[i], &records[i])
			if err != nil {
				panic(err)
			}
		}
	}
	return records
}

func (this CouchDb) Delete(record map[string]interface{}) bool {
	rev, err := this.Conn.Retrieve(fmt.Sprintf("%s",record["_id"]), &record)
	if err == nil && rev != "" {
		err = this.Conn.Delete(fmt.Sprintf("%s",record["_id"]), rev)
		if err == nil {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (this CouchDb) Update(record map[string]interface{}) bool {
	var r map[string]interface{}
	rev, err := this.Conn.Retrieve(fmt.Sprintf("%s",record["_id"]), &r)
	if err == nil {
		rev, err = this.Conn.EditWith(record, fmt.Sprintf("%s",record["_id"]), rev)
		if err == nil {
			return true
		} else {
			return false
		}
	} else {
		panic(err)
	}
}

