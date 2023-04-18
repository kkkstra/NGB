package middleware

import (
	"NGB/internal/config"
	"NGB/internal/controller/response"
	"NGB/internal/model"
	"NGB/pkg/jwt"
	"NGB/pkg/logrus"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	// skipper path, 支持正则表达式
	return func(c *gin.Context) {
		for _, skipPath := range config.C.User.Jwt.SkipPaths {
			// matched, err := regexp.MatchString(skipPath[1], c.FullPath())
			// if err != nil {
			// 	logrus.Logger.Error(err)
			// 	c.Abort()
			// 	return
			// }
			if c.Request.Method == skipPath[0] && c.FullPath() == skipPath[1] {
				logrus.Logger.Debugf("Skipped path: [%s]%s [%s]%s", skipPath[0], skipPath[1], c.Request.Method, c.FullPath())
				c.Next()
				return
			}
		}
		// get jwt
		bearerToken := c.Request.Header.Get("Authorization")
		bearerTokenSplit := strings.Split(bearerToken, " ")
		if len(bearerTokenSplit) < 2 {
			response.Error(c, http.StatusUnauthorized, "Empty token! ")
			c.Abort()
			return
		}
		token := bearerTokenSplit[1]
		if token == "" {
			response.Error(c, http.StatusUnauthorized, "Empty token! ")
			c.Abort()
			return
		}

		// parse jwt
		claims, err := jwt.ParseJWTToken(token)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Invalid token! ", err.Error())
			c.Abort()
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			response.Error(c, http.StatusUnauthorized, "Expired token! ")
			c.Abort()
			return
		}

		// 检查token是否因更改密码或删除用户而失效
		id := claims.Subject
		m := model.GetModel()
		u, err := m.FindUserById(id)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "User does not exist. ", err.Error())
			c.Abort()
			return
		}
		if u.UpdatePasswordAt.Unix() > claims.IssuedAt {
			response.Error(c, http.StatusUnauthorized, "Invalid toknen. ")
			c.Abort()
			return
		}

		userData := map[string]string{
			"id":   claims.Subject,
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
