package abstract

import "db"

type Database interface {
	Save(db.Patient) bool
	Read() []db.Patient
	Update(db.Patient) bool
	Delete(db.Patient) bool
}
