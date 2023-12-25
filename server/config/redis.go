package config

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // redis地址
	Password string `mapstructure:"password" json:"password" yaml:"password"` // redis密码
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis使用哪个数据库
}
