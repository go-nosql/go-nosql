package main

import (
	"db"
	"fmt"
)

func main() {
	database,_ := db.GetConnection("config.ini") //Pass configuration file location

	//Save
	//m["name"] = "suriya"
	//m["age"] = 24
	//m["city"] = "Cbe"
	//m["address"] = "India"
	//database.Save(m)

	//Read
	a := database.Read()
	fmt.Println(a)

	//Delete
	//database.Delete(a[0])

	//Update
	a[0]["city"] = "tirupur"
	database.Update(a[0])

}
