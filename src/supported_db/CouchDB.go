package supported_db

import (
	"couch-go-master"
	"db/entity"
)

type CouchDb struct {
	Conn couch.Database
}

func (this CouchDb) Save(patient entity.Patient) bool {
	id, rev, err := this.Conn.Insert(patient)
	if err == nil && id != "" && rev != "" {
		return true
	} else {
		return false
	}
}

func (this CouchDb) Read() []entity.Patient {
	ids, err := this.Conn.QueryIds("_all_docs", nil)
	patients := make([]entity.Patient, len(ids))
	if err != nil {
		panic(err)
	} else {
		for i := 0; i < len(ids); i++ {
			_, err = this.Conn.Retrieve(ids[i], &patients[i])
			if err != nil {
				panic(err)
			}
		}
	}
	return patients
}

func (this CouchDb) Delete(patient entity.Patient) bool {
	rev, err := this.Conn.Retrieve(patient.Id, &patient)
	if err == nil && rev != "" {
		err = this.Conn.Delete(patient.Id, rev)
		if err == nil {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (this CouchDb) Update(patient entity.Patient) bool {
	var r entity.Patient
	rev, err := this.Conn.Retrieve(patient.Id, &r)
	if err == nil {
		rev, err = this.Conn.EditWith(patient, patient.Id, rev)
		if err == nil {
			return true
		} else {
			return false
		}
	} else {
		panic(err)
	}
}

