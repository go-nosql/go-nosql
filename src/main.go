package main

import (
	"db"
	"fmt"
)

//main - main method to implement GO ORM - NoSQL
func main() {
	database := db.GetConnection("config.ini") //Pass configuration file location
	//record := db.NewObject()
	

	//Saving object
	//record.Set("student.name", "suriya")
	//record.Set("student.mark", 52)
	//record.Set("student.age", 24)
	//record.Set("employee.name", "williams")
	//record.Set("fruits",[]string{"apple","orange"})
	//database.Save(record)

	//Saving Json
	//jsn := db.NewJson()
	//jsn = `{"name":"hello"}`
	//database.Save(jsn.ToObject())
	

	//Read
	a := database.Read()
	//a := database.FindById("98c7c841105ee099229b90f0f7000318")
	//a := database.First()
	//a := database.Limit(2)
	//a := database.Count()
	//a := database.Where("mark ==  36")
	fmt.Println(a)

	//Converting to Json
	fmt.Println(a[0].ToJson())
	
	//Update
	//a[0].Set("student.communication.telephone","2332348")
	//database.Update(a[0])

	//Delete
	//database.Delete(a[0])
}
