package entity

type Detail struct {
	FirstName string `json:"FirstName" bson:"FirstName"`
	LastName  string `json:"LastName" bson:"LastName"`
	Dob       string `json:"Dob" bson:"Dob"`
}

type Contact struct {
	TelePhone string `json:"TelePhone" bson:"TelePhone"`
	Address   string `json:"Address" bson:"Address"`
}

type Patient struct {
	Id             string  `json:"_id" bson:"_id,omitempty"`
	PersonalDetail Detail  `json:"PersonalDetail" bson:"PersonalDetail"`
	ContactDetail  Contact `json:"ContactDetail" bson:"ContactDetail"`
	Height         float64 `json:"Height" bson:"Height"`
	Weight         float64 `json:"Weight" bson:"Weight"`
}
