package middleware

import (
	"hextech/common/response"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.FailWithDetailed(gin.H{"reload": true}, "服务器错误", c)
			}
		}()
		c.Next()
	}
}
