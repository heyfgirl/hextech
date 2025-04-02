package middleware

import (
	"hextech/common/response"
	"hextech/utils/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth(j *jwt.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}

		// 检查Bearer前缀
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.FailWithDetailed(gin.H{"reload": true}, "认证格式有误", c)
			c.Abort()
			return
		}

		token := parts[1]
		// 解析JWT token
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == jwt.ErrTokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		// 将用户信息存入上下文
		c.Set("claims", claims)
		c.Next()
	}
}
