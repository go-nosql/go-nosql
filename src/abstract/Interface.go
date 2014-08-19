package abstract

type Database interface {
	Save(interface{}) bool
	Read() []map[string]interface{}
	Update(map[string]interface{}) bool
	Delete(map[string]interface{}) bool
}
