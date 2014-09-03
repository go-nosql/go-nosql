package main

import (
	"github.com/go-nosql/go-nosql/src/db"
	"fmt"
)

//main - main method to implement GO ORM - NoSQL
func main() {

	//Return DB connection object based on config file
	database := db.GetConnection("config.ini") //Pass configuration file location

	//Create and return a new object
	//record := db.NewObject()

	//Saving object
	//record.Set("student.name", "suriya")
	//record.Set("student.mark", 52)
	//record.Set("student.age", 24)
	//record.Set("employee.name", "williams")
	//record.Set("fruits",[]string{"apple","orange"})
	//database.Save(record)

	//Saving Json
	//database.Save(`{"name":"hello"}`)

	//Read all records from database
	allRec := database.Read()
	fmt.Println(allRec)

	//Read record by id from database
	recById := database.FindById("53fd9d1eb0985415f02f75e2")
	fmt.Println(recById)

	//Read first record from database
	recFirst := database.First()
	fmt.Println(recFirst)
	
	//Read limited number of records from database
	recLimit := database.Limit(2)
	fmt.Println(recLimit)

	//Get record count from database
	count := database.Count()
	fmt.Println(count)
	
	//Get records based on query string.
	recWhere := database.Where("mark ==  36")
	fmt.Println(recWhere)

	//Converting to Json
	fmt.Println(allRec[0].ToJson())
	
	//Update using object
	//allRec[0].Set("student.communication.telephone","2332348")
	//database.Update(allRec[0])

	//Update using json
	//database.Update(`{"_id":"53fd9cf3b0985415f02f75e1","name":"world"}`)

	//Delete using object
	//database.Delete(allRec[1])
	
	//Delete using json
	//database.Delete(`{"_id":"53fd78d9b0985415f02f75de"}`)
}
