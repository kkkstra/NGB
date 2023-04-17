package controller

import (
	"NGB/internal/controller/param"
	"NGB/internal/controller/response"
	"NGB/internal/model"
	"NGB/pkg/jwt"
	"NGB/pkg/util"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var err error
	var req param.ReqSignUp
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind json! ", err.Error())
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to hash password! ", err.Error())
		return
	}
	m := model.GetModel()

	_, err = m.FindUserByUsername(req.Username)
	if err == nil {
		response.Error(c, http.StatusBadRequest, "Username already exists. ")
		return
	}

	u := &model.User{
		Username:         req.Username,
		Email:            req.Email,
		Password:         hashedPassword,
		Intro:            req.Intro,
		UpdatePasswordAt: time.Now(),
	}
	id, err := m.CreateUser(u)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create user! ", err.Error())
		return
	}
	response.Success(c, http.StatusOK, gin.H{"id": id}, "Sign up successfully! ")
}

func SignIn(c *gin.Context) {
	var req param.ReqSignIn
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind json! ", err.Error())
		return
	}
	m := model.GetModel()
	u, err := m.FindUserByUsername(req.Username)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "User does not exist. ", err.Error())
	}

	if ok, err := util.CheckPasswordHash(req.Password, u.Password); !ok {
		response.Error(c, http.StatusBadRequest, "Password is wrong! ", err.Error())
		return
	}

	//token, err := jwt.GetToken(json.Username, u.Role)
	tokenClaims := jwt.GenerateJWTToken(req.Username, u.Role, strconv.Itoa(int(u.ID)))
	token, err := jwt.GetJWTTokenString(tokenClaims)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get token! ", err.Error())
		return
	}
	response.Success(c, http.StatusOK, gin.H{"token": token, "expires_at": jwt.GetExpiresAt(tokenClaims)}, "Sign in successfully! ")
}

func GetUserProfile(c *gin.Context) {
	username := c.Param("username")
	m := model.GetModel()
	u, err := m.FindUserByUsername(username)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "User does not exist. ", err.Error())
		return
	}
	userData := gin.H{
		"username": u.Username,
		"email":    u.Email,
		"role":     u.Role,
		"intro":    u.Intro,
	}
	response.Success(c, http.StatusOK, userData, "Get user profile successfully. ")
}

func EditUserProfile(c *gin.Context) {
	if id, ok := checkAuthorization(c); ok {
		var req param.ReqEditProfile
		if err := c.ShouldBindJSON(&req); err != nil {
			response.Error(c, http.StatusBadRequest, "Failed to bind json! ", err.Error())
			return
		}

		m := model.GetModel()

		// 检查用户名是否重复
		_, err := m.FindUserByUsername(req.Username)
		if err == nil {
			response.Error(c, http.StatusBadRequest, "Username already exists. ")
			return
		}

		u := &model.User{
			Username: req.Username,
			Intro:    req.Intro,
		}
		err = m.UpdateUser(id, u)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed to update user profile! ", err.Error())
			return
		}
		response.Success(c, http.StatusCreated, gin.H{}, "Update user profile successfully! ")
	}
}

func EditUserPassword(c *gin.Context) {
	if id, ok := checkAuthorization(c); ok {
		var req param.ReqEditPassword
		if err := c.ShouldBindJSON(&req); err != nil {
			response.Error(c, http.StatusBadRequest, "Failed to bind json! ", err.Error())
			return
		}

		m := model.GetModel()
		u, err := m.FindUserById(id)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "User does not exist. ", err.Error())
		}

		if ok, err := util.CheckPasswordHash(req.OldPassword, u.Password); !ok {
			response.Error(c, http.StatusBadRequest, "Original password is wrong! ", err.Error())
			return
		}
		hashedPassword, err := util.HashPassword(req.NewPassword)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed to hash new password! ", err.Error())
			return
		}
		newUser := &model.User{
			Password:         hashedPassword,
			UpdatePasswordAt: time.Now(),
		}
		err = m.UpdateUser(id, newUser)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed to update user password! ", err.Error())
			return
		}
		response.Success(c, http.StatusCreated, gin.H{}, "Update user password successfully! ")
	}
}

// TODO
// 增加邮箱验证
func EditUserEmail(c *gin.Context) {
	if id, ok := checkAuthorization(c); ok {
		var req param.ReqEditEmail
		if err := c.ShouldBindJSON(&req); err != nil {
			response.Error(c, http.StatusBadRequest, "Failed to bind json! ", err.Error())
			return
		}

		m := model.GetModel()
		newUser := &model.User{
			Email: req.Email,
		}
		err := m.UpdateUser(id, newUser)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed to update user email! ", err.Error())
			return
		}
		response.Success(c, http.StatusCreated, gin.H{}, "Update user email successfully! ")
	}
}

func DeleteUser(c *gin.Context) {
	if id, ok := checkAuthorization(c); ok {
		m := model.GetModel()
		if err := m.DelUser(id); err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed to delete user. ", err.Error())
		}
		response.Success(c, http.StatusOK, gin.H{}, "Delete user successfully! ")
	}
}

func checkAuthorization(c *gin.Context) (string, bool) {
	username := c.Param("username")
	t, _ := c.Get("userdata")
	userData := t.(map[string]string)

	// get user id
	m := model.GetModel()

	u, err := m.FindUserByUsername(username)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "User does not exist. ", err.Error())
		return "", false
	}
	id := strconv.Itoa(int(u.ID))

	if userData["role"] != "admin" && userData["id"] != id {
		response.Error(c, http.StatusUnauthorized, "Insufficient permission. ")
		return "", false
	}

	return id, true
}
