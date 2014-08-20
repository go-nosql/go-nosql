go-orm
======

ORM with No-SQL

Deployment
==========

    1 . Download the package
    2 . Extract the file and paste it in to $GOROOT/src/pkg/
    3 . cd $GOROOT/src/pkg/go-orm-noschema/src/db/
    4 . go install (It will create $GOROOT/pkg/linux_amd64/go-orm-noschema/src/db.a file)

Example 
=======
package main

import (
            "go-orm-noschema/src/db"
            "fmt"
)

func main() {
            database, _ := db.GetConnection("config.ini") //pass configuration file location here
       	    a := database.Read()
            fmt.Println(a)
}
