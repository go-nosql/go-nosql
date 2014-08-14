package main

import (
	"db"
	"db/entity"
	"fmt"
)

func main() {
	database := db.GetConnection("config.ini") //Pass configuration file location

	//Save
	//patient := entity.Patient{}
	//patient.PersonalDetail.FirstName = "NewName"
	//database.Save(patient)

	//Read

	var patients []entity.Patient
	patients = database.Read()
	fmt.Println(patients)

	//Delete
	//database.Delete(patients[0])

	//Update
	//patients[0].PersonalDetail.FirstName = "updated"
	//database.Update(patients[0])
}
