package main

import (
	"fmt"
	"hextech/initialize"
)

func main() {
	// 初始化 Swagger
	initialize.InitSwagger()

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
	r := initialize.Routers(ctrl)

	// 启动服务
	r.Run(fmt.Sprintf(":%d", conf.System.Port))
}
