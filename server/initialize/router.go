package initialize

import (
	"server/global"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	// 设置为发布模式
	if global.GVA_Config.System.Env == "public" {
		gin.SetMode(gin.ReleaseMode) // Gin设为发布模式，确保在生产环境中减少冗余的日志和错误信息输出
	}
	Router := gin.New()

	// 安装插件
	InstallPlugin(Router)

	return Router

}
