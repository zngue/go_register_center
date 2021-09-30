package register

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/zngue/go_helper/pkg/api"
	"github.com/zngue/go_helper/pkg/code"
	"github.com/zngue/go_register_center/app/model"
	"github.com/zngue/go_register_center/app/service"
	"gorm.io/gorm"
	"time"
)

func ApiBase() service.IRegisterCenter {
	return service.NewRegisterCenter()
}

// Create  添加数据
func Create(ctx *gin.Context) {
	var req service.RegisterCenterRequest
	var data model.RegisterCenter
	if err := ctx.ShouldBind(&data); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	req.Data = &data
	var DelReq service.RegisterCenterRequest
	DelReq.Name = data.Name
	detailOne, err2 := ApiBase().Detail(DelReq)
	data.UpdateTime = time.Now().Unix()
	if err2 == gorm.ErrRecordNotFound {
		data.AddTime = time.Now().Unix()
		err := ApiBase().Add(req)
		api.DataWithErr(ctx, err, req.Data)
		return
	} else if err2 != nil {
		api.DataWithErr(ctx, err2, nil)
		return
	}
	if detailOne != nil && detailOne.ID > 0 {

		updateData := make(map[string]interface{})
		updateData["name"] = data.Name
		updateData["title"] = data.Title
		updateData["port"] = data.Port
		updateData["host"] = data.Host
		var updateReq service.RegisterCenterRequest
		updateReq.Name = data.Name
		updateReq.Data = updateData
		err3 := ApiBase().Save(updateReq)
		api.DataWithErr(ctx, err3, nil)
		return
	} else {
		data.AddTime = time.Now().Unix()
		err := ApiBase().Add(req)
		api.DataWithErr(ctx, err, req.Data)
		return
	}

}

// Edit  修改数据
func Edit(ctx *gin.Context) {
	var req service.RegisterCenterRequest
	var data model.RegisterCenter
	if err := ctx.ShouldBind(&data); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	req.ID = data.ID
	var ReqData map[string]interface{}

	if err := ctx.ShouldBind(&ReqData); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	//这里组装更新的数据  map[string]interface{}   数据和数据库的字段对应就行
	var updateData map[string]interface{}
	if ReqData != nil {
		file := new(code.FileNameChange)
		for key, val := range ReqData {
			mpKey := file.Camel2Case(file.Ucfirst(key))
			updateData[mpKey] = val
		}
	}
	//判断where 条件id是否存在
	if id, ok := updateData["id"]; ok {
		newID := cast.ToInt(id)
		if newID <= 0 {
			api.Error(ctx, api.Msg("id 不能为空"))
			return
		}
		req.ID = newID
	}
	//To do where 更新
	req.Data = updateData
	err := ApiBase().Save(req)
	api.DataWithErr(ctx, err, nil)
}

func List(ctx *gin.Context) {
	var req service.RegisterCenterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	list, err := ApiBase().List(req)
	api.DataWithErr(ctx, err, list)
}
func Detail(ctx *gin.Context) {
	var req service.RegisterCenterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	list, err := ApiBase().Detail(req)
	api.DataWithErr(ctx, err, list)
}
func Delete(ctx *gin.Context) {
	var req service.RegisterCenterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	err := ApiBase().Delete(req)
	api.DataWithErr(ctx, err, nil)
}
