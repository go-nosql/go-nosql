package main

import (
	"db"
	"db/entity"
	"fmt"
	_ "supported_db"
)

func main() {
	couch, mongo := db.GetConnection("config.ini") //Pass configuration file location
	fmt.Println("eeee : ", couch, "Mongo : ", mongo)

	//Save
	patient := entity.Patient{}
	patient.PersonalDetail.FirstName = "NewName"
	//mongo.Save(patient)
	couch.Save(patient)

	//Read

	var patients []entity.Patient
	//patients = mongo.Read()
	patients = couch.Read()
	fmt.Println(patients)

	//Delete

	//	mongo.Delete(patients[0])
	couch.Delete(patients[0])

	//Update

	patients[1].PersonalDetail.FirstName = "upgraded"
	//mongo.Update(patients[1])
	couch.Update(patients[1])
}
