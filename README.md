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
		* x2j 

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
        	database := db.GetConnection("config.ini") //Pass configuration file location
        	record := db.NewRecord() // Before calling NewRecord method GetConnection method should be called

        	//Save
        	record.Set("student.name", "suriya")
        	record.Set("student.mark", 52)
        	record.Set("student.age", 24)
        	record.Set("employee.name", "williams")
        	database.Save(record)

        	//Read
        	a := database.Read()
        	fmt.Println(a)

        	//Update
        	//a[0].Set("student.communication.telephone","2332348")
        	//database.Update(a[0])

        	//Delete
        	//database.Delete(a[0])
	}
