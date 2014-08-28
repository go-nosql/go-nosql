package supported_db

import (
	"testing"
	"abstract"
	"gopkg.in/mgo.v2/bson"
	"configparser-master"
	"gopkg.in/mgo.v2"
	"supported_db"
)

var database abstract.Database

// init - Establish database connection
func init() {
        conf, _ := configparser.Read("/home/visolve/go-orm/config.ini")
        db_config, _ := conf.Section("mongo")
        var mongoSession *mgo.Session
        if db_config.ValueOf("user") == "" {
                mongoSession, _ = mgo.Dial(db_config.ValueOf("ipaddress") + ":" + db_config.ValueOf("port"))
        } else {
                mongoSession, _ = mgo.Dial(db_config.ValueOf("user") + ":" + db_config.ValueOf("password") + "@" + db_config.ValueOf("ipaddress") + ":" + db_config.ValueOf("port") + "/" + db_config.ValueOf("dbname"))
        }
        db := mongoSession.DB(db_config.ValueOf("dbname"))
        collection := db.C(db_config.ValueOf("collectionname"))
        database = abstract.Database(supported_db.MongoDb{collection})
}

// TestRead - Test reading all the records
func TestRead(t *testing.T) {
	if database.Count() != len(database.Read()) {
		t.Fatalf("Error when reading records")
	}
}

// TestSave - Test saving a record
func TestSave(t *testing.T) {
	if database.Save(`{"name":"hello world"}`) == false {
		t.Fatalf("Error when saving record")
	}
}

// TestUpdate - Test updating a record
func TestUpdate(t *testing.T) {
	if database.Update(database.Read()[0].Set("name","abc")) == false {
		t.Fatalf("Error when updating record")
	}
}

// TestFirst - Test reading first record
func TestFirst(t *testing.T) {
	if database.Count() > 0 && database.First() == nil {
		t.Fatalf("Error when reading first record")
	}
}

// TestLast - Test reading last record
func TestLast(t *testing.T) {
        if database.Count() > 0 && database.Last() == nil {
                t.Fatalf("Error when reading last record")
        }
}

// TestCount - Test reading records count
func TestCount(t *testing.T) {
	if database.Count() < 0 {
		t.Fatalf("Error when reading records count")
	}
}

// TestLimit - Test reading limited records
func TestLimit(t *testing.T) {
	if database.Count() >= 1 && len(database.Limit(1)) != 1 {
		t.Fatalf("Error when reading limited records")
	}
}

// TestWhere - Test reading records with specified values
func TestWhere(t *testing.T) {
	if database.Count() > 0 {
		idString := database.Read()[0].Get("_id").(bson.ObjectId).Hex()
		if len(database.Where("_id == '" + idString + "'")) < 1 {
			t.Fatalf("Error when reading records with specified values")
		}
	}
}

// TestFindById - Test reading a record by its id
func TestFindById(t *testing.T) {
	if database.Count() > 0 {
		idString := database.Read()[0].Get("_id").(bson.ObjectId).Hex()
		if database.FindById(idString) == nil {
			t.Fatalf("Error when reading record with specified id")
		}
	}
}

// TestDelete - Test deleting a record
func TestDelete(t *testing.T) {
	if database.Delete(database.Read()[0]) == false {
		t.Fatalf("Error when deleting record")
	}
}
