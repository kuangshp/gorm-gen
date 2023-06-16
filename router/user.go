package router

import (
	"github.com/gin-gonic/gin"
	"gorm_gen/api"
	"gorm_gen/utils"
)

func InitUserRouter(Router *gin.RouterGroup) {
	pilotsRouter := Router.Group("user")
	pilotsRouter.GET("", api.GetAllUser)
	pilotsRouter.GET("/delete", api.DeleteUser)
	pilotsRouter.GET("/create", api.CreateUser)

	registerRouter := Router.Group("relation")
	newRegister := api.NewUser(&utils.DB)
	//fmt.Println(registerRouter, newRegister)
	registerRouter.GET("target", newRegister.GetAllUserApi) // 获取当前在什么群组
}
