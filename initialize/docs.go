package initialize

import (
	"hextech/docs"
)

// @title           DeepSeek API
// @version         1.0
// @description     DeepSeek 服务API文档
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey  Bearer
// @in                         header
// @name                       Authorization
// @description               Bearer token authentication

func InitSwagger() {
	// 设置 Swagger 信息
	docs.SwaggerInfo.Title = "DeepSeek API"
	docs.SwaggerInfo.Description = "DeepSeek 服务API文档"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
