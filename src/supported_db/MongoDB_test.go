package supported_db

import (
	"testing"
	"abstract"
	"gopkg.in/mgo.v2/bson"
	"github.com/alyu/configparser"
	"gopkg.in/mgo.v2"
)

var mdatabase abstract.Database

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
        mdatabase = abstract.Database(MongoDb{collection})
}

// TestRead - Test reading all the records
func TestMongoRead(t *testing.T) {
	if mdatabase.Count() != len(mdatabase.Read()) {
		t.Fatalf("Error when reading records")
	}
}

// TestSave - Test saving a record
func TestMongoSave(t *testing.T) {
	if mdatabase.Save(`{"name":"hello world"}`) == false {
		t.Fatalf("Error when saving record")
	}
	if mdatabase.Save(nil) != false {
		t.Fatalf("Error when saving record")
	}
}

// TestUpdate - Test updating a record
func TestMongoUpdate(t *testing.T) {
	if mdatabase.Update(mdatabase.Read()[mdatabase.Count()-1].Set("name","abc")) == false {
		t.Fatalf("Error when updating record")
	}
	if mdatabase.Update(nil) != false {
		t.Fatalf("Error when updating record")
	}
}

// TestFirst - Test reading first record
func TestMongoFirst(t *testing.T) {
	if mdatabase.Count() > 0 && mdatabase.First() == nil {
		t.Fatalf("Error when reading first record")
	}
}

// TestLast - Test reading last record
func TestMongoLast(t *testing.T) {
        if mdatabase.Count() > 0 && mdatabase.Last() == nil {
                t.Fatalf("Error when reading last record")
        }
}

// TestCount - Test reading records count
func TestMongoCount(t *testing.T) {
	if mdatabase.Count() < 0 {
		t.Fatalf("Error when reading records count")
	}
}

// TestLimit - Test reading limited records
func TestMongoLimit(t *testing.T) {
	if mdatabase.Count() >= 1 && len(mdatabase.Limit(1)) != 1 {
		t.Fatalf("Error when reading limited records")
	}
	if len(mdatabase.Limit(0)) != 0 {
		t.Fatalf("Error when reading limited records")
	}
}

// TestWhere - Test reading records with specified values
func TestMongoWhere(t *testing.T) {
	if mdatabase.Count() > 0 {
		idString := mdatabase.Read()[0].Get("_id").(bson.ObjectId).Hex()
		if len(mdatabase.Where("_id == '" + idString + "'")) < 1 {
			t.Fatalf("Error when reading records with specified values")
		}
                if len(mdatabase.Where("name == '" + mdatabase.Read()[mdatabase.Count()-1].Get("name").(string) + "'")) < 1 {
                        t.Fatalf("Error when reading records with specified values")
                }
	}
}

// TestFindById - Test reading a record by its id
func TestMongoFindById(t *testing.T) {
	if mdatabase.Count() > 0 {
		idString := mdatabase.Read()[0].Get("_id").(bson.ObjectId).Hex()
		if mdatabase.FindById(idString) == nil {
			t.Fatalf("Error when reading record with specified id")
		}
	}
	if mdatabase.FindById("invalidid") != nil {
		t.Fatalf("Error when reading record with specified id")
	}
}

// TestDelete - Test deleting a record
func TestMongoDelete(t *testing.T) {
	if mdatabase.Delete(mdatabase.Read()[mdatabase.Count()-1]) == false {
		t.Fatalf("Error when deleting record")
	}
	if mdatabase.Delete(nil) != false {
		t.Fatalf("Error when deleting record")
	}
}
