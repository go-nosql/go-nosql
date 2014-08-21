package entity

import "strings"

type Map map[string]interface{}


func (d Map) Get(keypath string) interface{} {
var segs []string = strings.Split(keypath, ".")
obj := d
for fieldIndex, field := range segs {
if fieldIndex == len(segs)-1 {
return obj[field]
}
switch obj[field].(type) {
case Map:
obj = obj[field].(Map)
case map[string]interface{}:
obj = Map(obj[field].(map[string]interface{}))
}
}
return obj
}


func (d Map) Set(keypath string, value interface{}) Map {
var segs []string
segs = strings.Split(keypath, ".")
obj := d
for fieldIndex, field := range segs {
if fieldIndex == len(segs)-1 {
obj[field] = value
}
if _, exists := obj[field]; !exists {
obj[field] = make(Map)
obj = obj[field].(Map)
} else {
switch obj[field].(type) {
case Map:
obj = obj[field].(Map)
case map[string]interface{}:
obj = Map(obj[field].(map[string]interface{}))
}
}
}
// chain
return d
}
