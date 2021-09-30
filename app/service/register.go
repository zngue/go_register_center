package service

import (
	"github.com/zngue/go_helper/pkg"
	"github.com/zngue/go_register_center/app/model"
)

type IRegisterCenter interface {
	Add(request RegisterCenterRequest) error
	Save(request RegisterCenterRequest) error
	List(request RegisterCenterRequest) (*[]model.RegisterCenter, error)
	Detail(request RegisterCenterRequest) (*model.RegisterCenter, error)
	Delete(request RegisterCenterRequest) error
	GetModel() interface{}
}
type RegisterCenter struct {
}
type RegisterCenterRequest struct {
	pkg.CommonRequest
	ID   int    `form:"id" field:"id" where:"eq" default:"0"`
	Name string `form:"name" field:"name" where:"eq" default:""`
}

func NewRegisterCenter() IRegisterCenter {
	return new(RegisterCenter)
}
func (i *RegisterCenter) GetModel() interface{} {
	return model.NewRegisterCenter()
}

// Add 添加
func (i *RegisterCenter) Add(request RegisterCenterRequest) error {
	request.ReturnType = 3
	return pkg.MysqlConn.Model(i.GetModel()).Create(request.Data).Error
}

// Save 修改
func (i *RegisterCenter) Save(request RegisterCenterRequest) error {
	request.ReturnType = 3
	return pkg.MysqlConn.Model(i.GetModel()).Updates(request.Data).Error
}
func (i *RegisterCenter) List(request RegisterCenterRequest) (*[]model.RegisterCenter, error) {
	var list []model.RegisterCenter
	err := request.Init(pkg.MysqlConn.Model(i.GetModel()), request).Find(&list).Error
	return &list, err
}

// Detail 详情
func (i *RegisterCenter) Detail(request RegisterCenterRequest) (*model.RegisterCenter, error) {
	var detail model.RegisterCenter
	request.ReturnType = 3
	err := request.Init(pkg.MysqlConn.Model(i.GetModel()), request).First(&detail).Error
	return &detail, err
}

// Delete 删除
func (i *RegisterCenter) Delete(request RegisterCenterRequest) error {
	request.ReturnType = 3
	return request.Init(pkg.MysqlConn, request).Delete(i.GetModel()).Error
}
