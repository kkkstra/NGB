package controller

import (
	"byitter/src/config"
	"byitter/src/controller/param"
	"byitter/src/controller/response"
	"byitter/src/model"
	"byitter/src/util"
	"byitter/src/util/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitAdmin() {
	username := config.C.User.Admin.Username
	password := config.C.User.Admin.Password
	email := config.C.User.Admin.Email
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		panic(err)
		return
	}
	m := model.GetModel(&model.UserModel{})

	if _, err = m.(*model.UserModel).FindUser(username); err != nil {
		u := &model.User{
			Username: username,
			Email:    email,
			Password: hashedPassword,
			Role:     1,
		}
		_, err := m.(*model.UserModel).CreateUser(u)
		if err != nil {
			panic(err)
			return
		}
		fmt.Println("Register admin successfully! ")
	} else {
		fmt.Println("Username already exists. ")
		return
	}
}

func SignUp(c *gin.Context) {
	var err error
	var json param.ReqSignUp
	if err := c.ShouldBindJSON(&json); err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind json! ", err.Error())
		return
	}
	hashedPassword, err := util.HashPassword(json.Password)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to hash password! ", err.Error())
		return
	}
	m := model.GetModel(&model.UserModel{})

	if _, err = m.(*model.UserModel).FindUser(json.Username); err != nil {
		u := &model.User{
			Username: json.Username,
			Email:    json.Email,
			Password: hashedPassword,
			Intro:    json.Intro,
			Github:   json.Github,
			School:   json.School,
			Website:  json.Website,
		}
		id, err := m.(*model.UserModel).CreateUser(u)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed to create user! ", err.Error())
			return
		}
		response.Success(c, gin.H{"id": id}, "Sign up successfully! ")
	} else {
		response.Error(c, http.StatusBadRequest, "Username already exists. ")
		return
	}
}

func SignIn(c *gin.Context) {
	var json param.ReqSignIn
	if err := c.ShouldBindJSON(&json); err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind json! ", err.Error())
		return
	}
	m := model.GetModel(&model.UserModel{})
	u, err := m.(*model.UserModel).FindUser(json.Username)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "User does not exist. ", err.Error())
	}

	if ok, err := util.CheckPasswordHash(json.Password, u.Password); ok {
		//token, err := jwt.GetToken(json.Username, u.Role)
		token := jwt.GenerateUserJwt(json.Username, u.Role)
		tokenStr, err := token.GenerateTokenStr()
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed to get token! ", err.Error())
			return
		}
		response.Success(c, gin.H{"token": tokenStr, "ExpiresAt": token.ExpiresAt()}, "Sign in successfully! ")
		return
	} else {
		response.Error(c, http.StatusBadRequest, "Password is wrong! ", err.Error())
		return
	}
}

func GetUserProfile(c *gin.Context) {
	username := c.Param("username")
	m := model.GetModel(&model.UserModel{})
	u, err := m.(*model.UserModel).FindUser(username)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "User does not exist. ", err.Error())
		return
	}
	userData := gin.H{
		username:  u.Username,
		"email":   u.Email,
		"role":    u.Role.Str(),
		"intro":   u.Intro,
		"github":  u.Github,
		"school":  u.School,
		"website": u.Website,
	}
	response.Success(c, userData, "Get user profile successfully. ")
}

func EditUserProfile(c *gin.Context) {
	username := c.Param("username")
	t, _ := c.Get("userdata")
	userData := t.(map[string]string)
	if userData["role"] != "admin" && userData["username"] != username {
		response.Error(c, http.StatusUnauthorized, "Insufficient permission. ")
		return
	}
	c.JSON(http.StatusOK, gin.H{"a": "A"})
}
