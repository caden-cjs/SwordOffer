/*
@Time : 2020/04/08 20:45
@Author : caden
@File : Singleton.go
@Software: GoLand
*/
package module

import (
	"reflect"
	"sync"
)

var singleton map[string]interface{}
var lock sync.Mutex
var once sync.Once
func init() {
	singleton = make(map[string]interface{})
}
func GetSingleton(typeT interface{}) interface{} {
	ttype := reflect.TypeOf(typeT)
	name := ttype.String()
	data := singleton[name]
	if data != nil {
		return data
	} else {
		//lock.Lock()
		//defer lock.Unlock()
		once.Do(func() {
			singleton[name] = typeT
		})
		//singleton[name] = typeT
		return typeT
	}

}
