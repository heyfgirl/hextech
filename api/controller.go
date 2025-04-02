package api

import (
	"hextech/common/response"
	"hextech/config"
	"hextech/utils/jwt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	JWT    *jwt.JWT
	DB     *gorm.DB
	Config *config.Config
}

// @Summary     健康检查
// @Description 检查服务是否正常运行
// @Tags        基础接口
// @Produce     json
// @Success     200 {object} map[string]interface{}
// @Router      /health [get]

func (ctrl *Controller) HealthCheck(c *gin.Context) {
	response.OkWithMessage("服务正常运行", c)
}

// @Summary     Ping测试
// @Description 测试服务是否可响应
// @Tags        基础接口
// @Produce     json
// @Success     200 {object} map[string]interface{}
// @Router      /ping [get]

func (ctrl *Controller) Ping(c *gin.Context) {
	response.OkWithMessage("pong", c)
}
