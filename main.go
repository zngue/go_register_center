package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg"
	"github.com/zngue/go_helper/pkg/api"
	"github.com/zngue/go_helper/pkg/common_run"
	"github.com/zngue/go_register_center/app/model"
	"github.com/zngue/go_register_center/app/router"
)

func main() {
	common_run.CommonGinRun(
		common_run.RedisLoad(true),
		common_run.ConfigLoad(true),
		common_run.MysqlLoad(true),

		common_run.FnRouter(func(engine *gin.Engine) {
			pkg.MysqlConn.AutoMigrate(new(model.RegisterCenter))
			engine.NoRoute(func(context *gin.Context) {
				api.Success(context, api.Code(500), api.Msg("page is not found"))
			})
			router.Router(engine.Group(""))
			engine.GET("ok", func(context *gin.Context) {
				api.Success(context,api.Msg("ok"))
			})
		}),
	)
}
