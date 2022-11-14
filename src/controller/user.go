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
	"strconv"
	"time"
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

	_, err = m.(*model.UserModel).FindUserByUsername(username)
	if err == nil {
		fmt.Println("Username already exists. ")
		return
	}

	u := &model.User{
		Username:         username,
		Email:            email,
		Password:         hashedPassword,
		Role:             1,
		UpdatePasswordAt: time.Now(),
	}
	_, err = m.(*model.UserModel).CreateUser(u)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("Register admin successfully! ")
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

	_, err = m.(*model.UserModel).FindUserByUsername(json.Username)
	if err == nil {
		response.Error(c, http.StatusBadRequest, "Username already exists. ")
		return
	}

	u := &model.User{
		Username:         json.Username,
		Email:            json.Email,
		Password:         hashedPassword,
		Intro:            json.Intro,
		Github:           json.Github,
		School:           json.School,
		Website:          json.Website,
		UpdatePasswordAt: time.Now(),
	}
	id, err := m.(*model.UserModel).CreateUser(u)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create user! ", err.Error())
		return
	}
	response.Success(c, http.StatusOK, gin.H{"id": id, "role": "common"}, "Sign up successfully! ")
}

func SignIn(c *gin.Context) {
	var json param.ReqSignIn
	if err := c.ShouldBindJSON(&json); err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind json! ", err.Error())
		return
	}
	m := model.GetModel(&model.UserModel{})
	u, err := m.(*model.UserModel).FindUserByUsername(json.Username)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "User does not exist. ", err.Error())
	}

	if ok, err := util.CheckPasswordHash(json.Password, u.Password); !ok {
		response.Error(c, http.StatusBadRequest, "Password is wrong! ", err.Error())
		return
	}

	//token, err := jwt.GetToken(json.Username, u.Role)
	token := jwt.GenerateUserJwt(json.Username, u.Role, strconv.Itoa(int(u.ID)))
	tokenStr, err := token.GenerateTokenStr()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get token! ", err.Error())
		return
	}
	response.Success(c, http.StatusOK, gin.H{"token": tokenStr, "expires_at": token.ExpiresAt()}, "Sign in successfully! ")
}

func GetUserProfile(c *gin.Context) {
	username := c.Param("username")
	m := model.GetModel(&model.UserModel{})
	u, err := m.(*model.UserModel).FindUserByUsername(username)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "User does not exist. ", err.Error())
		return
	}
	userData := gin.H{
		"username": u.Username,
		"email":    u.Email,
		"role":     u.Role.Str(),
		"intro":    u.Intro,
		"github":   u.Github,
		"school":   u.School,
		"website":  u.Website,
	}
	response.Success(c, http.StatusOK, userData, "Get user profile successfully. ")
}

func EditUserProfile(c *gin.Context) {
	username := c.Param("username")
	t, _ := c.Get("userdata")
	userData := t.(map[string]string)
	if userData["role"] != "admin" && userData["username"] != username {
		response.Error(c, http.StatusUnauthorized, "Insufficient permission. ")
		return
	}
	m := model.GetModel(&model.UserModel{})
	var id string
	if userData["role"] != "admin" {
		id = userData["id"]
	} else {
		u, err := m.(*model.UserModel).FindUserByUsername(username)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "User does not exist. ", err.Error())
			return
		}
		id = strconv.Itoa(int(u.ID))
	}

	var json param.ReqEditProfile
	if err := c.ShouldBindJSON(&json); err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind json! ", err.Error())
		return
	}

	u := &model.User{
		Intro:   json.Intro,
		Github:  json.Github,
		School:  json.School,
		Website: json.Website,
	}
	err := m.(*model.UserModel).UpdateUser(id, u)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update user profile! ", err.Error())
		return
	}
	response.Success(c, http.StatusCreated, gin.H{}, "Update user profile successfully! ")
}

func EditUserPassword(c *gin.Context) {
	username := c.Param("username")
	t, _ := c.Get("userdata")
	userData := t.(map[string]string)
	if userData["role"] != "admin" && userData["username"] != username {
		response.Error(c, http.StatusUnauthorized, "Insufficient permission. ")
		return
	}
	m := model.GetModel(&model.UserModel{})
	var id string
	if userData["role"] != "admin" {
		id = userData["id"]
	} else {
		u, err := m.(*model.UserModel).FindUserByUsername(username)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "User does not exist. ", err.Error())
			return
		}
		id = strconv.Itoa(int(u.ID))
	}

	var json param.ReqEditPassword
	if err := c.ShouldBindJSON(&json); err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind json! ", err.Error())
		return
	}

	u, err := m.(*model.UserModel).FindUserById(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "User does not exist. ", err.Error())
	}

	if ok, err := util.CheckPasswordHash(json.OldPassword, u.Password); !ok {
		response.Error(c, http.StatusBadRequest, "Original password is wrong! ", err.Error())
		return
	}
	hashedPassword, err := util.HashPassword(json.NewPassword)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to hash new password! ", err.Error())
		return
	}
	newUser := &model.User{
		Password:         hashedPassword,
		UpdatePasswordAt: time.Now(),
	}
	err = m.(*model.UserModel).UpdateUser(id, newUser)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update user password! ", err.Error())
		return
	}
	response.Success(c, http.StatusCreated, gin.H{}, "Update user password successfully! ")
}

func EditUserEmail(c *gin.Context) {
	username := c.Param("username")
	t, _ := c.Get("userdata")
	userData := t.(map[string]string)
	if userData["role"] != "admin" && userData["username"] != username {
		response.Error(c, http.StatusUnauthorized, "Insufficient permission. ")
		return
	}
	m := model.GetModel(&model.UserModel{})
	var id string
	if userData["role"] != "admin" {
		id = userData["id"]
	} else {
		u, err := m.(*model.UserModel).FindUserByUsername(username)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "User does not exist. ", err.Error())
			return
		}
		id = strconv.Itoa(int(u.ID))
	}

	var json param.ReqEditEmail
	if err := c.ShouldBindJSON(&json); err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind json! ", err.Error())
		return
	}

	newUser := &model.User{
		Email: json.Email,
	}
	err := m.(*model.UserModel).UpdateUser(id, newUser)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update user email! ", err.Error())
		return
	}
	response.Success(c, http.StatusCreated, gin.H{}, "Update user email successfully! ")
}

func DeleteUser(c *gin.Context) {
	username := c.Param("username")
	t, _ := c.Get("userdata")
	userData := t.(map[string]string)
	if userData["role"] != "admin" && userData["username"] != username {
		response.Error(c, http.StatusUnauthorized, "Insufficient permission. ")
		return
	}
	m := model.GetModel(&model.UserModel{})
	var id string
	if userData["role"] != "admin" {
		id = userData["id"]
	} else {
		u, err := m.(*model.UserModel).FindUserByUsername(username)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "User does not exist. ", err.Error())
			return
		}
		id = strconv.Itoa(int(u.ID))
	}
	if err := m.(*model.UserModel).DelUser(id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete user. ", err.Error())
	}
	response.Success(c, http.StatusOK, gin.H{}, "Delete user successfully! ")
}
