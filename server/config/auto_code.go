package config

type AutoCode struct {
	Root            string `mapstructure:"root" json:"root" yaml:"root"`                                        // 自动生成代码的根目录
	Server          string `mapstructure:"server" json:"server" yaml:"server"`                                  // 自动生成代码的server目录
	SPlug           string `mapstructure:"server-plug" json:"server-plug" yaml:"server-plug"`                   // 自动生成代码的
	SInitialize     string `mapstructure:"server-initialize" json:"server-initialize" yaml:"server-initialize"` // 自动生成代码的server
	SModel          string `mapstructure:"server-model" json:"server-model" yaml:"server-model"`                // 自动生成代码的
	SRequest        string `mapstructure:"server-request" json:"server-request" yaml:"server-request"`          // 自动生成代码的
	SRouter         string `mapstructure:"server-router" json:"server-router" yaml:"server-router"`             // 自动生成代码的
	SService        string `mapstructure:"server-service" json:"server-service" yaml:"server-service"`          // 自动生成代码的
	SApi            string `mapstructure:"server-api" json:"server-api" yaml:"server-api"`                      // 自动生成代码的
	Web             string `mapstructure:"web" json:"web" yaml:"web"`                                           // 自动生成代码的web目录
	WAPI            string `mapstructure:"web-api" json:"web-api" yaml:"web-api"`                               // 自动生成代码的web
	WForm           string `mapstructure:"web-form" json:"web-form" yaml:"web-form"`                            // 自动生成代码的
	WTable          string `mapstructure:"web-table" json:"web-table" yaml:"web-table"`                         // 自动生成代码的
	TransferRestart bool   `mapstructure:"transfer-restart" json:"transfer-restart" yaml:"transfer-restart"`    // 是否重启自动生成
}
