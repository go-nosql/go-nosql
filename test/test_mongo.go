package main

import (
	"supported_db"
	"entity"
	"fmt"
	)

func main(){

	dbObj := supported_db.MongoDb{}

	//Save
	/*
	patient := entity.Patient{}
	patient.PersonalDetail.FirstName = "initial"
	dbObj.Save(patient)
	*/

	//Read
	var patients []entity.Patient
	patients = dbObj.Read()

	//Delete
	/*
	fmt.Println(patients)
	dbObj.Delete(patients[1])
	*/

	//Update
	patients[1].PersonalDetail.FirstName = "upgraded"
	fmt.Println(dbObj.Update(patients[1]))
}
