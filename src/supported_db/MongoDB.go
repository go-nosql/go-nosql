package supported_db 

import (
    "mgo"
    "mgo/bson"
    "entity"
    "configparser-master"
)

type MongoDb struct{
}

func (this *MongoDb) GetConnection() *mgo.Collection{
		var collection *mgo.Collection
                conf, err := configparser.Read("/home/visolve/go-orm/config.ini")
                if err != nil {
			panic(err)
                }
                confg,_ := conf.Section("database")
		mongoSession, err := mgo.Dial(confg.ValueOf("IPAddress")+":"+confg.ValueOf("Port"))
                if err != nil {
			panic(err)
                }else{
			db := mongoSession.DB(confg.ValueOf("DbName"))
			collection = db.C(confg.ValueOf("CollectionName"))
		}
		return collection
}



func (this *MongoDb) Read() []entity.Patient{
            var patients []entity.Patient
            this.GetConnection().Find(nil).All(&patients)
	    return patients
}


func (this *MongoDb) Save(patient entity.Patient) bool{
        err := this.GetConnection().Insert(patient)
	if(err==nil){
		return true
	}else{
		return false
	}
}



func (this *MongoDb) Delete(patient entity.Patient) bool{
                err := this.GetConnection().Remove(bson.M{"_id" : bson.ObjectId(patient.Id)})
		if(err==nil){
			return true
		}else{
			return false
		}
}


func (this *MongoDb) Update(patient entity.Patient) bool{
	err := this.GetConnection().UpdateId(bson.ObjectId(patient.Id),bson.M{"PersonalDetail" : patient.PersonalDetail, "ContactDetail" : patient.ContactDetail,"Height":patient.Height,"Weight":patient.Weight})
	if(err==nil){
		return true
	}else{
		return false
	}
}
