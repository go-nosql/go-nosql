package abstract

import "github.com/go-nosql/go-nosql/src/db/entity"

// Database - Interface to define basic DB functionalities.
type Database interface {
	Save(interface{}) bool
	Read() []entity.Map
	Update(interface{}) bool
	Delete(interface{}) bool
        First() entity.Map
        Last() entity.Map
	Count() int
	Limit(int) []entity.Map
	FindById(string) entity.Map
	Where(string) []entity.Map
	Merge(interface{}) bool
}
