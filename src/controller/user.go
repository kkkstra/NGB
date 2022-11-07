package controller

import (
	"byitter/src/controller/param"
	"byitter/src/controller/response"
	"byitter/src/model"
	"byitter/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
	m := model.GetModel()

	if _, err = m.FindUser(json.Username); err != nil {
		u := &model.UserModel{
			Username: json.Username,
			Email:    json.Email,
			Password: hashedPassword,
			Intro:    json.Intro,
			Github:   json.Github,
			School:   json.School,
			Website:  json.Website,
		}
		id, err := m.CreateUser(u)
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
	m := model.GetModel()
	u, err := m.FindUser(json.Username)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "User does not exist. ", err.Error())
	}

	if ok, err := util.CheckPasswordHash(json.Password, u.Password); ok {
		response.Success(c, gin.H{"id": u.ID}, "Sign in successfully! ")
		return
	} else {
		response.Error(c, http.StatusBadRequest, "Password is wrong! ", err.Error())
		return
	}
}
