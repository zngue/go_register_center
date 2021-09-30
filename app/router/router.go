package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_register_center/app/api/register"
)

func Router(group *gin.RouterGroup) {
	group.POST("register", register.Create)
}
func registerRouters(group *gin.RouterGroup) {
	registerRouter := group.Group("register")
	{
		registerRouter.POST("create", register.Create)
		registerRouter.POST("edit", register.Edit)
		registerRouter.POST("delete", register.Delete)
		registerRouter.GET("list", register.List)
		registerRouter.GET("detail", register.Detail)
	}
}
