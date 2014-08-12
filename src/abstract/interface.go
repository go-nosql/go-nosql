package abstract

import "entity"

type Database interface {
	//GetConnection()
	Save(entity.Patient) bool
	Read() []entity.Patient
	Update(entity.Patient) bool
	Delete(entity.Patient) bool
}
