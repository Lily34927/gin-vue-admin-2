package config

type GeneralDB struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                         // 数据库地址
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                         // 数据库端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                   // 数据库配置
	DbName       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                // 数据库名称
	Password     string `mapstructure:"password" json:"password" yaml:"password"`             // 数据库密码
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`             // 日志模式
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle"` // 空闲连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open"` // 最大连接数
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                // 日志模式
}
