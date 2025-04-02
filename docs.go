// Package main 主程序包
package main

import (
	// 加载服务，启动swagger
	_ "hextech/docs"
)

// @title API服务接口文档
// @version v1.0
// @description 这是API服务的Swagger文档

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

func init() {
	enableSwagger = true
}

// 启动 swagger 需要 go run 同时启动 docs.go 和 main.go
// 例如 go run docs.go main.go
