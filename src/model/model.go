package model

import (
	"reflect"
)

type ModelInterface interface{}

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
