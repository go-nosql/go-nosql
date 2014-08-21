package abstract

import "db/entity"

type Database interface {
	Save(interface{}) bool
	Read() []entity.Map
	Update(map[string]interface{}) bool
	Delete(map[string]interface{}) bool
}
