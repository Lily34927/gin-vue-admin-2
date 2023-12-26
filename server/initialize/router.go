package initialize

import (
	"net/http"
	"server/global"
	"server/router"

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
	systemRouter := router.RouterGroupAPP.System
	exampleRouter := router.RouterGroupAPP.Example

	Router.StaticFS(global.GVA_Config.Local.StorePath, http.Dir(global.GVA_Config.Local.StorePath))

	// swagger

	PublicGroup := Router.Group(global.GVA_Config.System.RouterPrefix)
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基本功能路由 不做鉴权
		systemRouter.InitInitRouter(PublicGroup) // 自动化相关功能
	}
	PrivateGroup := Router.Group(global.GVA_Config.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup, PublicGroup)
		systemRouter.InitJwtRouter(PrivateGroup)
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
		systemRouter.InitSysRouter(PrivateGroup)
		systemRouter.InitCasbinRouter(PrivateGroup)
		systemRouter.InitAutoCodeRouter(PrivateGroup)
		systemRouter.InitAuthorityRouter(PrivateGroup)
		systemRouter.InitSysDictionaryRouter(PrivateGroup)
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)
		systemRouter.InitAuthorityBtnRouter(PrivateGroup)
		systemRouter.InitChatGptRouter(PrivateGroup)

		exampleRouter.InitCustomerRouter(PrivateGroup)
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup)
	}

	global.GVA_LOG.Info("router register success")
	return Router
}
