package main

import (
	"supported_db"
	"entity"
	_"fmt"
	)

func main(){

	dbObj := supported_db.MongoDb{}
	//dbObj := supported_db.CouchDb{}

	//Save
	patient := entity.Patient{}
	patient.PersonalDetail.FirstName = "first"
	dbObj.Save(patient)

	//Read
	/*
	var patients []entity.Patient
	patients = dbObj.Read()
	fmt.Println(patients)
	*/

	//Delete
	/*
	dbObj.Delete(patients[0])
	*/

	//Update
	/*
	patients[0].PersonalDetail.FirstName = "upgraded"
	dbObj.Update(patients[0])
	*/
}
