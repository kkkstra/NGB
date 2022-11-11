package initEnv

import (
	"byitter/src/controller"
	"byitter/src/util/jwt"
)

func InitEnv() {
	jwt.InitRSAKey()
	controller.InitAdmin()
}
