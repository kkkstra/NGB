package initenv

import (
	"byitter/src/config"
	"byitter/src/controller"
	"byitter/src/util/jwt"
)

func InitEnv() {
	jwt.InitRSAKey()
	if config.C.Init.Admin {
		controller.InitAdmin()
	}
}
