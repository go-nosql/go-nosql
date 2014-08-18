package supported_db

import (
	"db/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type MongoDb struct {
	Conn *mgo.Collection
}

func (this MongoDb) Read() []entity.Patient {
	var patients []entity.Patient
	this.Conn.Find(nil).All(&patients)
	for i, _ := range patients {
		patients[i].Id = fmt.Sprintf("%x", string(patients[i].Id))
	}
	return patients
}

func (this MongoDb) Save(patient entity.Patient) bool {
	err := this.Conn.Insert(patient)
	if err == nil {
		return true
	} else {
		return false
	}
}

func (this MongoDb) Delete(patient entity.Patient) bool {
	err := this.Conn.Remove(bson.M{"_id": bson.ObjectIdHex(patient.Id)})
	if err == nil {
		return true
	} else {
		return false
	}
}

func (this MongoDb) Update(patient entity.Patient) bool {
	err := this.Conn.UpdateId(bson.ObjectIdHex(patient.Id), bson.M{"PersonalDetail": patient.PersonalDetail, "ContactDetail": patient.ContactDetail, "Height": patient.Height, "Weight": patient.Weight})
	if err == nil {
		return true
	} else {
		return false
	}
}
