package initialize

import (
	"hextech/api"
	"hextech/config"
	"hextech/utils/jwt"

	"gorm.io/gorm"
)

func InitController(config *config.Config, jwt *jwt.JWT, db *gorm.DB) *api.Controller {
	return &api.Controller{
		Config: config,
		JWT:    jwt,
		DB:     db,
	}
}
