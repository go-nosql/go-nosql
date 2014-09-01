package supported_db

import (
	"testing"
	"abstract"
	"configparser-master"
	"couch-go-master"
)

var cdatabase abstract.Database

// init - Establish database connection
func init() {
	//database = db.GetConnection("/home/visolve/go-orm/config.ini")
	conf, _ := configparser.Read("/home/visolve/go-orm/config.ini")
	db_config, _ := conf.Section("couch")
        var db couch.Database
        if db_config.ValueOf("user") == "" {
                db, _ = couch.NewDatabase(db_config.ValueOf("ipaddress"), db_config.ValueOf("port"), db_config.ValueOf("dbname"))
        } else {
                db, _ = couch.NewDatabaseByURL("http://" + db_config.ValueOf("user") + ":" + db_config.ValueOf("password") + "@" + db_config.ValueOf("ipaddress") + ":" + db_config.ValueOf("port") + "/" + db_config.ValueOf("dbname"))
        }
        cdatabase = abstract.Database(CouchDb{db})
}

// TestRead - Test reading all the records
func TestCouchRead(t *testing.T) {
	if cdatabase.Count() != len(cdatabase.Read()) {
		t.Fatalf("Error when reading records")
	}
}

// TestSave - Test saving a record
func TestCouchSave(t *testing.T) {
	if cdatabase.Save(`{"name":"hello world"}`) == false {
		t.Fatalf("Error when saving record")
	}
	if cdatabase.Save(nil) != false {
		t.Fatalf("Error when saving record")
	}
}

// TestUpdate - Test updating a record
func TestCouchUpdate(t *testing.T) {
	if cdatabase.Update(cdatabase.Read()[cdatabase.Count()-1].Set("name","hello abc")) == false {
		t.Fatalf("Error when updating record")
	}
	if cdatabase.Update(nil) != false {
		t.Fatalf("Error when updating record")
	}
}

// TestFirst - Test reading first record
func TestCouchFirst(t *testing.T) {
	if cdatabase.Count() > 0 && cdatabase.First() == nil {
		t.Fatalf("Error when reading first record")
	}
}

// TestLast - Test reading last record
func TestCouchLast(t *testing.T) {
        if cdatabase.Count() > 0 && cdatabase.Last() == nil {
                t.Fatalf("Error when reading last record")
        }
}

// TestCount - Test reading records count
func TestCouchCount(t *testing.T) {
	if cdatabase.Count() < 0 {
		t.Fatalf("Error when reading records count")
	}
}

// TestLimit - Test reading limited records
func TestCouchLimit(t *testing.T) {
	if cdatabase.Count() >= 1 && len(cdatabase.Limit(1)) != 1 {
		t.Fatalf("Error when reading limited records")
	}
	if len(cdatabase.Limit(-5)) != 0 {
		t.Fatalf("Error when reading limited records")
	}
}

// TestWhere - Test reading records with specified values
func TestCouchWhere(t *testing.T) {
	if cdatabase.Count() > 0 {
		idString := cdatabase.Read()[0].Get("_id").(string)
		if len(cdatabase.Where("_id == '" + idString + "'")) < 1 {
			t.Fatalf("Error when reading records with specified values")
		}
		if len(cdatabase.Where("name == '" + cdatabase.Read()[cdatabase.Count()-1].Get("name").(string) + "'")) < 1 {
			t.Fatalf("Error when reading records with specified values")
		}
	}
}

// TestFindById - Test reading a record by its id
func TestCouchFindById(t *testing.T) {
	if cdatabase.Count() > 0 {
		idString := cdatabase.Read()[0].Get("_id").(string)
		if cdatabase.FindById(idString) == nil {
			t.Fatalf("Error when reading record with specified id")
		}
	}
        if cdatabase.FindById("invalidid") != nil {
                t.Fatalf("Error when reading record with specified id")
        }
}

// TestDelete - Test deleting a record
func TestCouchDelete(t *testing.T) {
	if cdatabase.Delete(cdatabase.Read()[cdatabase.Count()-1]) == false {
		t.Fatalf("Error when deleting record")
	}
	if cdatabase.Delete(nil) != false {
		t.Fatalf("Error when deleting record")
	}
}
