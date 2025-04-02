package config

type Config struct {
	System System `mapstructure:"system" json:"system"`
	MySQL  MySQL  `mapstructure:"mysql" json:"mysql"`
	JWT    JWT    `mapstructure:"jwt" json:"jwt"`
}

type System struct {
	Port         int    `mapstructure:"port" json:"port"`
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
}

type MySQL struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         int    `mapstructure:"port" json:"port"`
	Config       string `mapstructure:"config" json:"config"`
	DbName       string `mapstructure:"db-name" json:"db-name"`
	Username     string `mapstructure:"username" json:"username"`
	Password     string `mapstructure:"password" json:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns"`
}

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signing-key"`   // jwt签名
	ExpiresTime int64  `mapstructure:"expires-time" json:"expires-time"` // 过期时间
	BufferTime  int64  `mapstructure:"buffer-time" json:"buffer-time"`   // 缓冲时间
	Issuer      string `mapstructure:"issuer" json:"issuer"`             // 签发者
}
