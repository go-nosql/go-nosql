package abstract

import "db/entity"

type Database interface {
	Save(entity.Patient) bool
	Read() []entity.Patient
	Update(entity.Patient) bool
	Delete(entity.Patient) bool
}
