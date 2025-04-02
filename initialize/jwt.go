package initialize

import (
	"hextech/config"
	"hextech/utils/jwt"
)

func InitJWT(cfg config.JWT) *jwt.JWT {
	return jwt.NewJWT(cfg.SigningKey, cfg.ExpiresTime, cfg.BufferTime, cfg.Issuer)
}
