package main

import (
	"supported_db"
	"entity"
	_"fmt"
	)

func main(){

	dbObj := supported_db.CouchDb{}

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
	dbObj.Delete(patients[0])
	*/

	//Update
	patients[0].PersonalDetail.FirstName = "upgraded"
	dbObj.Update(patients[0])
}
