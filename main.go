package main

import (
	"fmt"
	"hextech/initialize"
	"hextech/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var enableSwagger = false

func main() {

	// 初始化配置
	conf := initialize.InitializeConfig()

	jwt := initialize.InitJWT(conf.JWT)
	fmt.Println(jwt)

	// 初始化数据库
	db := initialize.GormMysql(conf.MySQL)
	if db == nil {
		panic("数据库连接失败")
	}

	// 初始化控制器
	ctrl := initialize.InitController(conf, jwt, db)

	// 初始化路由
	r := gin.New()
	r.Use(middleware.DefaultLogger())
	r.Use(middleware.RecoveryMiddleware())

	// 启动swagger 需要 go run 同时启动 docs.go 和 main.go
	if enableSwagger {
		r.GET("/swagger", func(c *gin.Context) {
			c.Redirect(301, "/swagger/index.html")
		})
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	r = initialize.Routers(r, ctrl)

	// 启动服务
	r.Run(fmt.Sprintf(":%d", conf.System.Port))
}
