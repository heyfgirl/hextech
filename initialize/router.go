package initialize

import (
	"hextech/api"
	"hextech/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API文档
// @version         1.0
// @description     这是一个示例API服务器。
// @host           localhost:8080
// @BasePath       /
func Routers(ctrl *api.Controller) *gin.Engine {
	r := gin.Default()

	// r.Use(middleware.DefaultLogger())

	// 注册Swagger路由
	r.GET("/swagger", func(c *gin.Context) {
		c.Redirect(301, "/swagger/index.html")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	publicGroup := r.Group(ctrl.Config.System.RouterPrefix)
	authGroup := r.Group(ctrl.Config.System.RouterPrefix)
	authGroup.Use(middleware.JWTAuth(ctrl.JWT))

	{
		// 注册基础路由组
		InitBaseRouter(ctrl, r.Group(""))
	}
	{
		// 注册用户相关路由组
		InitUserRouter(ctrl, publicGroup, authGroup)
	}

	return r
}

// 基础服务
func InitBaseRouter(ctrl *api.Controller, Router *gin.RouterGroup) {
	baseRouter := Router.Group("")
	{
		baseRouter.GET("/health", ctrl.HealthCheck)
		baseRouter.GET("/ping", ctrl.Ping)
	}
}

func InitUserRouter(ctrl *api.Controller, publicGroup *gin.RouterGroup, authGroup *gin.RouterGroup) {
	notAuthUserGroup := publicGroup.Group("/api/v1/user")
	{
		notAuthUserGroup.POST("/register", ctrl.Register)
		notAuthUserGroup.POST("/login", ctrl.Login)
	}
	authUserGroup := authGroup.Group("/api/v1/user")
	{
		authUserGroup.GET("/info", ctrl.GetUserInfo)
	}
}
