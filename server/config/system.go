package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                  // 环境
	DbType        string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                      // 数据库类型
	OssType       string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`                   // OSS类型
	RouterPrefix  string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`    // 路由前缀
	Addr          int    `mapstructure:"intaddr" json:"addr" yaml:"addr"`                            // 监听地址
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`    // IP限制次数
	LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`       // IP限制时间
	UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                // 是否使用Redis
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` // 是否使用多点登录
}
