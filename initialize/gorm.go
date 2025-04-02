package initialize

import (
	"fmt"
	"hextech/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GormMysql(mysqlConfig config.MySQL) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		mysqlConfig.Username,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DbName,
		mysqlConfig.Config,
	)

	mysqlConf := mysql.Config{
		DSN:                      dsn,
		DefaultStringSize:        191,
		DisableDatetimePrecision: true,
	}
	if db, err := gorm.Open(mysql.New(mysqlConf), &gorm.Config{}); err != nil {
		panic(fmt.Errorf("MySQL启动异常: %s", err))
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConns)
		return db
	}
}
