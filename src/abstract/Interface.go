package abstract

import "db/entity"

// Database - Interface to define basic DB functions.
type Database interface {
	Save(interface{}) bool
	Read() []entity.Map
	Update(map[string]interface{}) bool
	Delete(map[string]interface{}) bool
}
