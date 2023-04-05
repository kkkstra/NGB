package model

import (
	"reflect"
)

type ModelInterface interface{}

// TODO
// 这里有无更好的写法？
func GetModel(Model interface{}) interface{} {
	t := reflect.TypeOf(Model)
	switch t.String() {
	case "*model.UserModel":
		return &UserModel{db}
	case "*model.RSAKeyModel":
		return &RSAKeyModel{db}
	}
	return nil
}
