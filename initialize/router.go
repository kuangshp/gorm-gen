package initialize

import (
	"github.com/gin-gonic/gin"
	"gorm_gen/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 设置获取ip地址
	Router.ForwardedByClientIP = true
	// 注册中间件
	//Router.Use()
	// 配置全局路径
	ApiGroup := Router.Group("/api/v1/admin")
	// 注册路由
	router.InitUserRouter(ApiGroup) // 账号中心
	router.InitMessageRouter(ApiGroup)
	return Router
}
