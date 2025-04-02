package main

import (
	"fmt"
	"hextech/initialize"

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

	ctrl := initialize.InitController(conf, jwt, db)
	// 初始化路由
	r := gin.Default()
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
