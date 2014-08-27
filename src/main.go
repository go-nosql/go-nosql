package main

import (
	"db"
	"fmt"
)

//main - main method to implement GO ORM - NoSQL
func main() {

	//Return DB connection object based on config file
	database := db.GetConnection("config.ini") //Pass configuration file location

	//Create and return a new object
	record := db.NewObject()

	//Saving object
	record.Set("student.name", "suriya")
	record.Set("student.mark", 52)
	record.Set("student.age", 24)
	record.Set("employee.name", "williams")
	record.Set("fruits",[]string{"apple","orange"})
	database.Save(record)

	//Saving Json
	database.Save(`{"name":"hello"}`)

	//Read all records from database
	a := database.Read()
	fmt.Println(a)

	//Read record by id from database
	a := database.FindById("98c7c841105ee099229b90f0f7000318")
	fmt.Println(a)

	//Read first record from database
	a := database.First()
	fmt.Println(a)
	
	//Read limited number of records from database
	a := database.Limit(2)
	fmt.Println(a)

	//Get record count from database
	a := database.Count()
	fmt.Println(a)
	
	//Get records based on query string.
	a := database.Where("mark ==  36")
	fmt.Println(a)

	//Converting to Json
	fmt.Println(a[0].ToJson())
	
	//Update using object
	a[0].Set("student.communication.telephone","2332348")
	database.Update(a[0])

	//Update using json
	database.Update(`{"_id":"4deaf29629ea5cf3438cb3043100397d","name":"hello"}`)

	//Delete using object
	database.Delete(a[0])
	
	//Delete using json
	database.Delete(`{"_id":"53fd78d9b0985415f02f75de"}`)
}
