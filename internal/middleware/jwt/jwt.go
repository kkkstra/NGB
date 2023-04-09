package jwt

import (
	"NGB/internal/controller/response"
	"NGB/internal/model"
	"NGB/pkg/util"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get jwt
		bearerToken := c.Request.Header.Get("Authorization")
		bearerTokenSplit := strings.Split(bearerToken, " ")
		if len(bearerTokenSplit) < 2 {
			response.Error(c, http.StatusUnauthorized, "Empty token! ")
			return
		}
		token := bearerTokenSplit[1]
		if token == "" {
			response.Error(c, http.StatusUnauthorized, "Empty token! ")
			return
		}

		// parse jwt
		claims, err := util.ParseJWTToken(token)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Invalid token! ", err.Error())
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			response.Error(c, http.StatusUnauthorized, "Expired token! ")
			return
		}

		// 检查token是否因更改密码或删除用户而失效
		id := claims.Id
		m := model.GetModel()
		u, err := m.FindUserById(id)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "User does not exist. ", err.Error())
			return
		}
		if u.UpdatePasswordAt.Unix() > claims.IssuedAt {
			response.Error(c, http.StatusUnauthorized, "Invalid toknen. ")
			return
		}

		userData := map[string]string{
			"id":   claims.Id,
			"role": getRoleString(claims.Role),
		}
		c.Set("userdata", userData)
		c.Next()
	}
}

func getRoleString(role int) string {
	if role == 1 {
		return "admin"
	}
	return "common"
}
