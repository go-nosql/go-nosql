go-orm
======

ORM for No-SQL databases

PreRequisites
=============

	* CouchDB
	* MongoDB
	* golang packages //Refer Deployment section to install golang packages
		* couch-go
		* mgo
		* configparser

CouchDB Installation [Linux]
============================

	yum install couchdb
	vim /etc/couchdb/local.ini
		[httpd]
		;port = 5984
		;bind_address = 127.0.0.1
		bind_address = 172.16.1.87 (Your IP address)
	service couchdb start

MongoDB Installation [Linux]
============================

	vim /etc/yum.repos.d/mongodb.repo                 
		[mongodb]
		name=MongoDB Repository
		baseurl=http://downloads-distro.mongodb.org/repo/redhat/os/x86_64/
		gpgcheck=0
		enabled=1
	yum install mongodb-org
        service mongod start

Deployment
==========

	1 . Download the package from github
	2 . Extract the file in $GOROOT/src/pkg/
	3 . cd $GOROOT/src/pkg/go-orm/src/db/
	4 . go install (It will create $GOROOT/pkg/linux_amd64/go-orm/src/db.a file)

Example 
=======

	package main

	import (
            "go-orm/src/db"
            "fmt"
	)

	func main() {

		//Return DB connection object based on config file
        	database := db.GetConnection("config.ini") //Pass configuration file location

		//Create and return a new object
        	record := db.NewRecord() // Before calling NewRecord method GetConnection method should be called

        	//Saving Object
        	record.Set("student.name", "suriya")
        	record.Set("student.mark", 52)
        	record.Set("student.age", 24)
        	record.Set("employee.name", "williams")
        	database.Save(record)

	        //Saving Json
	        //database.Save(`{"name":"hello"}`)

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

		/*
		 Where clause supports single condition only
		 String comparison : database.Where("name == 'williams'") Must specify single quotes
		 Numeric comparison : database.Where("student.mark > 52") 
		*/
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
