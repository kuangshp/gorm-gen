package router

import (
	"github.com/gin-gonic/gin"
	"gorm_gen/api"
	"gorm_gen/utils"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	registerRouter := Router.Group("message")
	newRegister := api.NewMessage(&utils.DB)
	registerRouter.GET("", newRegister.GetMessageListApi) //
	registerRouter.POST("", newRegister.CreateMessageApi)
}
