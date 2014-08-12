package supported_db 

import (
    "couch-go-master"
    "configparser-master"
    "entity"
)

type CouchDb struct{
}

func (this *CouchDb) GetConnection() couch.Database{
                conf, err := configparser.Read("/home/visolve/go-orm/config.ini")
                if err != nil {
			panic(err)
                }
                confg,_ := conf.Section("database")
                db,err := couch.NewDatabase(confg.ValueOf("IPAddress"), confg.ValueOf("Port"), confg.ValueOf("DbName"))
                if err != nil {
			panic(err)
                }
		return db
}

func (this *CouchDb) Save(patient entity.Patient) bool{
        id, rev, err := this.GetConnection().Insert(patient)
        if(err==nil&&id!=""&&rev!=""){
			return true
                }else{
                        return false
                }
}


func (this *CouchDb) Read() []entity.Patient{
                ids,err := this.GetConnection().QueryIds("_all_docs",nil)
		patients := make([]entity.Patient, len(ids))
                if(err!=nil){
			panic(err)
                }else{
                        for i:=0;i<len(ids);i++{
                                _, err = this.GetConnection().Retrieve(ids[i], &patients[i])
				if(err!=nil){
					panic(err)
				}
                        }
                }
		return patients
}


func (this *CouchDb) Delete(patient entity.Patient) bool{
                rev,err := this.GetConnection().Retrieve(patient.Id,&patient)
                if(err==nil&&rev!=""){
                        err = this.GetConnection().Delete(patient.Id,rev)
                        if(err==nil){
				return true
                        }else{
				return false
                        }
                }else{
			return false
                }
}

func (this *CouchDb) Update(patient entity.Patient) bool{
		var r entity.Patient
		rev,err := this.GetConnection().Retrieve(patient.Id,&r)
                if(err==nil){
                        	rev,err = this.GetConnection().EditWith(patient,patient.Id,rev)
				if(err==nil){
					return true
				}else{
					return false
				}
                        }else{
				panic(err)
                }
}
