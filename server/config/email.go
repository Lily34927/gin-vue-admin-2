package config

type Email struct {
	To       string `mapstructure:"to" json:"to" yaml:"to"`                   // 收件人
	From     string `mapstructure:"from" json:"from" yaml:"from"`             // 发件人
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 邮箱服务器
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`       // 授权码
	NickName string `mapstructure:"nickname" json:"nickname" yaml:"nickname"` // 昵称

	Port  int  `mapstructure:"port" json:"port" yaml:"port"`       // 端口
	IsSSL bool `mapstructure:"is-ssl" json:"is-ssl" yaml:"is-ssl"` // 是否SSL
}
