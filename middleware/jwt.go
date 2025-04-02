package middleware

import (
	"errors"
	"hextech/common/request"
	"hextech/common/response"
	"hextech/utils/jwt"

	"github.com/gin-gonic/gin"
)

func JWTAuth(j *jwt.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}

		// 解析JWT token
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		// 将用户信息存入上下文
		request.SetClaims(c, claims)
		c.Next()
	}
}
