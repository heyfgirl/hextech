package initialize

import (
	"fmt"
	"hextech/config"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func InitializeConfig() *config.Config {
	configType := getConfigType()
	v := viper.New()

	// 设置配置文件名称（不带扩展名）
	v.SetConfigName("config")

	// 支持的配置文件类型
	v.SetConfigType(configType)

	// 配置文件路径
	v.AddConfigPath(".")        // 当前目录
	v.AddConfigPath("./config") // config目录

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到
			panic(fmt.Sprintf("配置文件未找到: %s", err))
		} else {
			// 其他错误
			panic(fmt.Sprintf("读取配置文件失败: %s", err))
		}
	}

	var conf config.Config
	if err := v.Unmarshal(&conf); err != nil {
		panic(fmt.Sprintf("配置文件解析失败: %s", err))
	}

	return &conf
}

// getConfigType 根据环境变量或文件存在情况决定使用哪种配置文件
func getConfigType() string {
	// 首先检查环境变量
	if configType := os.Getenv("CONFIG_TYPE"); configType != "" {
		return strings.ToLower(configType)
	}

	// 检查配置文件是否存在
	if fileExists("config.yaml") || fileExists("config.yml") {
		return "yaml"
	}
	if fileExists("config.json") {
		return "json"
	}

	// 默认使用 yaml
	return "yaml"
}

// fileExists 检查文件是否存在
func fileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}
