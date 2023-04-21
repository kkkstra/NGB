package controller

import (
	"NGB/internal/config"
	"NGB/internal/controller/param"
	"NGB/internal/controller/response"
	"NGB/internal/model"
	"NGB/internal/model/redis"
	"NGB/pkg/jwt"
	"NGB/pkg/util"
	"fmt"
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

	// 检查用户名是否存在
	_, err = m.FindUserByUsername(req.Username)
	if err == nil {
		response.Error(c, http.StatusBadRequest, "Username already exists. ")
		return
	}
	// 检查邮箱是否重复
	_, err = m.FindUserByEmail(req.Email)
	if err == nil {
		response.Error(c, http.StatusBadRequest, "Email already exists. ")
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
	response.Success(c, http.StatusOK, response.Data{"id": id}, "Sign up successfully! ")
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

	if req.Method == "password" {
		if ok, err := util.CheckPasswordHash(req.Password, u.Password); !ok {
			response.Error(c, http.StatusBadRequest, "Password is wrong! ", err.Error())
			return
		}
	} else if req.Method == "code" {
		cli := redis.GetClient()
		code, err := cli.GetCode(u.Email)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed to get code. ", err.Error())
			return
		}
		if code != req.Code {
			response.Success(c, http.StatusInternalServerError, response.Data{}, "Wrong code. ")
			return
		}
	} else {
		response.Error(c, http.StatusBadRequest, "Wrong method. ", "")
		return
	}

	//token, err := jwt.GetToken(json.Username, u.Role)
	tokenClaims := jwt.GenerateJWTToken(req.Username, u.Role, strconv.Itoa(int(u.ID)), config.C.User.Jwt.Expire, config.C.User.Jwt.Issuer)
	token, err := jwt.GetJWTTokenString(tokenClaims)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get token! ", err.Error())
		return
	}
	response.Success(c, http.StatusOK, response.Data{"token": token, "expires_at": jwt.GetExpiresAt(tokenClaims)}, "Sign in successfully! ")
}

func GetSignInCode(c *gin.Context) {
	var req param.ReqGetSignInCode
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind json! ", err.Error())
		return
	}
	m := model.GetModel()
	_, err := m.FindUserByEmail(req.Email)
	if err != nil {
		response.Success(c, http.StatusOK, response.Data{}, "Sending email complete! ")
		return
	}
	// 生成验证码
	code := util.GenerateValidateCode(6)
	cli := redis.GetClient()
	// 检测邮件发送频率
	sendTime, err := cli.GetSendMailTime(req.Email)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to send email. ", err.Error())
		return
	}
	nowTime := time.Now().Unix()
	if sendTime != -1 && nowTime-sendTime < int64(config.C.User.Code.MailFrequency*60) {
		response.Success(c, http.StatusOK, response.Data{}, "Send mail too frequently. ")
		return
	}
	// 存储验证码
	cli.UpdateCode(req.Email, code)
	// 发送邮件
	err = util.SendEmail(util.CustomEmail{
		From:    fmt.Sprintf("%s <%s>", config.C.Email.Sender, config.C.Email.Account),
		To:      req.Email,
		Subject: "NGB: Validate Code",
		Content: fmt.Sprintf("Your validate code is %s. ", code),
		Account: config.C.Email.Account,
		Code:    config.C.Email.Code,
		Addr:    config.C.Email.Addr,
		Server:  config.C.Email.Server,
	})
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to send email", err.Error())
		return
	}
	err = cli.UpdateSendMailTime(req.Email, time.Now().Unix())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update send time. ", "")
		return
	}
	response.Success(c, http.StatusOK, response.Data{}, "Send mail successfully. ")
}

func GetUserProfile(c *gin.Context) {
	username := c.Param("username")
	m := model.GetModel()
	u, err := m.FindUserByUsername(username)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "User does not exist. ", err.Error())
		return
	}
	userData := response.Data{
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
		if req.Username != c.Param("username") {
			_, err := m.FindUserByUsername(req.Username)
			if err == nil {
				response.Error(c, http.StatusBadRequest, "Username already exists. ")
				return
			}
		}

		u := &model.User{
			Username: req.Username,
			Intro:    req.Intro,
		}
		err := m.UpdateUser(id, u)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed to update user profile! ", err.Error())
			return
		}
		response.Success(c, http.StatusCreated, response.Data{}, "Update user profile successfully! ")
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
		response.Success(c, http.StatusCreated, response.Data{}, "Update user password successfully! ")
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

		// 检查邮箱是否重复
		_, err := m.FindUserByEmail(req.Email)
		if err == nil {
			response.Error(c, http.StatusBadRequest, "Email already exists. ")
			return
		}

		newUser := &model.User{
			Email: req.Email,
		}
		err = m.UpdateUser(id, newUser)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed to update user email! ", err.Error())
			return
		}
		response.Success(c, http.StatusCreated, response.Data{}, "Update user email successfully! ")
	}
}

func DeleteUser(c *gin.Context) {
	if id, ok := checkAuthorization(c); ok {
		m := model.GetModel()
		if err := m.DelUser(id); err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed to delete user. ", err.Error())
		}
		response.Success(c, http.StatusOK, response.Data{}, "Delete user successfully! ")
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
