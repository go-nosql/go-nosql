package supported_db

import (
	"github.com/go-nosql/go-nosql/src/db/entity"
	"strings"
	"strconv"
	"os/exec"
	"reflect"
	"fmt"
	"encoding/json"
)

// GtmDb - Struct for GT.M database.
type GtmDb struct {
	Conn string
	MFilePath string
}

// Read - Read all records from GT.M database.
func (this GtmDb) Read() []entity.Map {
	var records []entity.Map
	out , err := exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" do read^GtmDb").Output()
   	if err == nil {
		var v interface{}
		json.Unmarshal(out, &v)
		record := entity.Map(v.(map[string]interface{}))
	        for k, v := range record {
			records = append(records,entity.Map{k:v})
		}
    	}
	return records
}

// Save - Save generic record in GT.M database.
func (this GtmDb) Save(record interface{}) bool {
	if record == nil {
		return false
	}
        if reflect.TypeOf(record).String() == "string" {
                record = entity.Json(record.(string)).ToObject() //convert Json to Map object
        }
	recordDirect := removeNestingGtm(map[string]interface{}(record.(entity.Map)),"")
	for k, v := range recordDirect {
		if strings.Contains(reflect.TypeOf(v).String(), "int") || strings.Contains(reflect.TypeOf(v).String(), "float") {
			_, _ = exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" set ^" + this.Conn + "(" + k + ")=" + fmt.Sprintf("%v",v)).Output()
		} else if strings.Contains(reflect.TypeOf(v).String(), "bool") {
			if v.(bool) {
				_, _ = exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" set ^" + this.Conn + "(" + k + ")=1").Output()
			} else {
				_, _ = exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" set ^" + this.Conn + "(" + k + ")=0").Output()
			}
		} else if strings.Contains(reflect.TypeOf(v).String(), "string") {
			_, _ = exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" set ^" + this.Conn + "(" + k + ")=\"" + v.(string) + "\"").Output()
		}
	}
	return true
}

// Delete - Delete generic record in GT.M database.
func (this GtmDb) Delete(record interface{}) bool {
	if record == nil {
		return false
	}
	for k, _ := range record.(entity.Map) {
	        out , err := exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" write $$delete^GtmDb(\"" + string(k) + "\")").Output()
        	if err == nil {
                        s := ""
                        for _, field := range out {
                                s = s+string(field)
                        }
			s=s[0:len(s)-1]
			if s=="1" {
				return true
			}
                }
	}
	return true
}

// Update - Update record in GT.M database.
func (this GtmDb) Update(record interface{}) bool {
        if record == nil {
                return false
        }
        if reflect.TypeOf(record).String() == "string" {
                record = entity.Json(record.(string)).ToObject() //convert Json to Map object
        }
        recordDirect := removeNestingGtm(map[string]interface{}(record.(entity.Map)),"")
	uniqKey := ""
        for k, v := range recordDirect {
		rootKey := strings.Split(k,",")[0]
		if uniqKey != rootKey {
			uniqKey = rootKey
	        	_ , _ = exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" write $$delete^GtmDb(" + uniqKey + ")").Output()
		}
                if strings.Contains(reflect.TypeOf(v).String(), "int") || strings.Contains(reflect.TypeOf(v).String(), "float") {
                        _, _ = exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" set ^" + this.Conn + "(" + k + ")=" + fmt.Sprintf("%v",v)).Output()
                } else if strings.Contains(reflect.TypeOf(v).String(), "bool") {
                        if v.(bool) {
                                _, _ = exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" set ^" + this.Conn + "(" + k + ")=1").Output()
                        } else {
                                _, _ = exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" set ^" + this.Conn + "(" + k + ")=0").Output()
                        }
                } else if strings.Contains(reflect.TypeOf(v).String(), "string") {
                        _, _ = exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" set ^" + this.Conn + "(" + k + ")=\"" + v.(string) + "\"").Output()
                }
        }
        return true
}

// First - Read first record from GT.M database
func (this GtmDb) First() entity.Map {
	var rec entity.Map
        out , err := exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" do first^GtmDb").Output()
        if err == nil {
                var v interface{}
                json.Unmarshal(out, &v)
                record := entity.Map(v.(map[string]interface{}))
                for k, v := range record {
                        rec = entity.Map{k:v}
                        }
                }
        return rec
}

// Last - Read last record from GT.M database
func (this GtmDb) Last() entity.Map {
        var rec entity.Map
        out , err := exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" do last^GtmDb").Output()
        if err == nil {
                var v interface{}
                json.Unmarshal(out, &v)
                record := entity.Map(v.(map[string]interface{}))
                for k, v := range record {
                        rec = entity.Map{k:v}
                        }
                }
        return rec
}

// Count - Read number of records from GT.M database
func (this GtmDb) Count() int {
        out , err := exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" do count^GtmDb").Output()
        if err == nil {
			s := ""
		        for _, field := range out {
		                s = s+string(field)
		        }
			s=s[0:len(s)-1]
			num, err1 := strconv.Atoi(s)
			if err1 == nil {
				return num
			}
                }
        return 0
}

// Limit - Read limited number of records from GT.M database
func (this GtmDb) Limit(limit int) []entity.Map {
        var records []entity.Map
        if limit <= 0 {
                return records
        }
        out , err := exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" do limit^GtmDb(" + strconv.Itoa(limit) + ")").Output()
        if err == nil {
                var v interface{}
                json.Unmarshal(out, &v)
                record := entity.Map(v.(map[string]interface{}))
                for k, v := range record {
                        records = append(records,entity.Map{k:v})
                        }
                }
        return records
}

// Where - Get the records based on query string.
func (this GtmDb) Where(query string) []entity.Map {
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
	s := strings.Join(all[0:2]," ")
	s = s + " " + fmt.Sprintf("%v",val)
        out , err := exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" do where^GtmDb(\"" + s + "\")").Output()
        if err == nil {
                var v interface{}
                json.Unmarshal(out, &v)
                record := entity.Map(entity.Map(v.(map[string]interface{})).Get("matchRec").(map[string]interface{}))
                for k, v := range record {
                        records = append(records,entity.Map{k:v})
                        }
                }
        return records
}

// FindById - Read record by id from GT.M database.
func (this GtmDb) FindById(id string) entity.Map {
        var rec entity.Map
        out , err := exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" do byId^GtmDb(\"" + id + "\")").Output()
        if err == nil && len(out) > 0 {
                var v interface{}
                json.Unmarshal(out, &v)
                record := entity.Map(v.(map[string]interface{}))
                for k, v := range record {
                        rec = entity.Map{k:v}
                        }
                }
        return rec
}

// Merge - Merge user given record and GT.M database record
func (this GtmDb) Merge(record interface{}) bool {
        if record == nil {
                return false
        }
        if reflect.TypeOf(record).String() == "string" {
                record = entity.Json(record.(string)).ToObject() //convert Json to Map object
        }
        recordDirect := removeNestingGtm(map[string]interface{}(record.(entity.Map)),"")
        for k, v := range recordDirect {
                if strings.Contains(reflect.TypeOf(v).String(), "int") || strings.Contains(reflect.TypeOf(v).String(), "float") {
                        _, _ = exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" set ^" + this.Conn + "(" + k + ")=" + fmt.Sprintf("%v",v)).Output()
                } else if strings.Contains(reflect.TypeOf(v).String(), "bool") {
                        if v.(bool) {
                                _, _ = exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" set ^" + this.Conn + "(" + k + ")=1").Output()
                        } else {
                                _, _ = exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" set ^" + this.Conn + "(" + k + ")=0").Output()
                        }
                } else if strings.Contains(reflect.TypeOf(v).String(), "string") {
                        _, _ = exec.Command("mumps","-run","%XCMD","zlink \"" + this.MFilePath + "\" set ^db=\"" + this.Conn + "\" set ^" + this.Conn + "(" + k + ")=\"" + v.(string) + "\"").Output()
                }
        }
        return true
}

// removeNesting - Remove nested maps to create single level map
func removeNestingGtm(src map[string]interface{}, chain string) map[string]interface{} {
        obj := make(map[string]interface{})
        if chain != "" {
                chain = chain + ","
                }
        for k, v := range src {
                        switch v.(type) {
                        case map[string]interface{}:
                                temp := make(map[string]interface{})
                                temp = removeNestingGtm(v.(map[string]interface{}), chain + "\"" + k + "\"")
                                for k1, v1 := range temp {
                                        obj[k1] = v1
                                }
                        default:
                                obj[chain + "\"" + k + "\""] = v
                        }
                }
        return obj
}
