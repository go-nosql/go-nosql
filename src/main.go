package main

import (
	"db"
	"db/entity"
	"fmt"
	_ "supported_db"
)

func main() {
	database := db.GetConnection("config.ini") //Pass configuration file location

	//Save
	patient := entity.Patient{}
	patient.PersonalDetail.FirstName = "NewName"
	//mongo.Save(patient)
	database.Save(patient)

	//Read

	var patients []entity.Patient
	//patients = mongo.Read()
	patients = database.Read()
	fmt.Println(patients)

	//Delete

	//	mongo.Delete(patients[0])
	database.Delete(patients[0])

	//Update

	patients[1].PersonalDetail.FirstName = "updated"
	//mongo.Update(patients[1])
	database.Update(patients[1])
}
