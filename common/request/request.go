package request

import (
	"hextech/utils/jwt"

	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) *jwt.CustomClaims {
	claims, ok := c.Get("claims")
	if !ok {
		panic("claims is not found")
	}
	return claims.(*jwt.CustomClaims)
}

func GetUserID(c *gin.Context) uint {
	claims := GetClaims(c)
	return claims.BaseClaims.ID
}

func GetUserName(c *gin.Context) string {
	claims := GetClaims(c)
	return claims.BaseClaims.Username
}

func SetClaims(c *gin.Context, claims *jwt.CustomClaims) {
	c.Set("claims", claims)
}
