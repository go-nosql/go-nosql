package supported_db

import (
	"testing"
	"github.com/go-nosql/go-nosql/src/abstract"
	"github.com/alyu/configparser"
)

var gdatabase abstract.Database

// init - Establish database connection
func init() {
        conf, _ := configparser.Read("/usr/local/go/src/github.com/go-nosql/go-nosql/config.ini")
        db_config, _ := conf.Section("gtm")
        gdatabase = abstract.Database(GtmDb{Conn:db_config.ValueOf("dbname"),MFilePath:db_config.ValueOf("mfilepath")})
}

// TestGtmRead - Test reading all the records
func TestGtmRead(t *testing.T) {
	if gdatabase.Count() != len(gdatabase.Read()) {
		t.Fatalf("Error when reading records")
	}
}

// TestGtmSave - Test saving a record
func TestGtmSave(t *testing.T) {
	if gdatabase.Save(`{"name":"hello world"}`) == false {
		t.Fatalf("Error when saving record")
	}
	if gdatabase.Save(nil) != false {
		t.Fatalf("Error when saving record")
	}
}

// TestGtmUpdate - Test updating a record
func TestGtmUpdate(t *testing.T) {
	if gdatabase.Update(gdatabase.Read()[gdatabase.Count()-1].Set("name","abc")) == false {
		t.Fatalf("Error when updating record")
	}
	if gdatabase.Update(nil) != false {
		t.Fatalf("Error when updating record")
	}
}

// TestGtmFirst - Test reading first record
func TestGtmFirst(t *testing.T) {
	if gdatabase.Count() > 0 && gdatabase.First() == nil {
		t.Fatalf("Error when reading first record")
	}
}

// TestGtmLast - Test reading last record
func TestGtmLast(t *testing.T) {
        if gdatabase.Count() > 0 && gdatabase.Last() == nil {
                t.Fatalf("Error when reading last record")
        }
}

// TestGtmCount - Test reading records count
func TestGtmCount(t *testing.T) {
	if gdatabase.Count() < 0 {
		t.Fatalf("Error when reading records count")
	}
}

// TestGtmLimit - Test reading limited records
func TestGtmLimit(t *testing.T) {
	if gdatabase.Count() >= 1 && len(gdatabase.Limit(1)) != 1 {
		t.Fatalf("Error when reading limited records")
	}
	if len(gdatabase.Limit(0)) != 0 {
		t.Fatalf("Error when reading limited records")
	}
}

// TestGtmWhere - Test reading records with specified values
func TestGtmWhere(t *testing.T) {
	if gdatabase.Count() > 0 {
		if len(gdatabase.Where("invalidId == invalidId")) != 0 {
			t.Fatalf("Error when reading records with specified values")
		}
	}
}

// TestGtmFindById - Test reading a record by its id
func TestGtmFindById(t *testing.T) {
	if gdatabase.FindById("invalidid") != nil {
		t.Fatalf("Error when reading record with specified id")
	}
}

// TestGtmMerge - Test merging user record and GtmDB record
func TestGtmMerge(t *testing.T) {
        if gdatabase.Count() > 0 {
                lastRec := gdatabase.Read()[gdatabase.Count()-1]
		lastRec.Set("newField","newValue")
                if gdatabase.Merge(lastRec) == false {
                        t.Fatalf("Error when merging user record and GtmDB record")
                }
        }
}

// TestGtmDelete - Test deleting a record
func TestGtmDelete(t *testing.T) {
	if gdatabase.Delete(gdatabase.Read()[gdatabase.Count()-1]) == false {
		t.Fatalf("Error when deleting record")
	}
	if gdatabase.Delete(nil) != false {
		t.Fatalf("Error when deleting record")
	}
}
