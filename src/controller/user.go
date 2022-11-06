package controller

import (
	"byitter/src/model"
	"byitter/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUpData struct {
	Username string `json:"username" binding:"required,min=6,max=32"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=64"`
	Intro    string `json:"intro" binding:"max=512"`
}

// SignUp 用户注册
func SignUp(c *gin.Context) {
	var json SignUpData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := util.HashPassword(json.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	m := model.GetModel()
	u := &model.User{
		Username: json.Username,
		Email:    json.Email,
		Password: hashedPassword,
		Intro:    json.Intro,
	}
	id, err := m.Insert(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id, "msg": "Sign up successfully! "})
}
