package service

import (
	"github.com/zngue/go_helper/pkg"
	"github.com/zngue/go_register_center/app/model"
)

type IUser interface {
	Add(request UserRequest) error
	Save(request UserRequest) error
	List(request UserRequest) (*[]model.User, error)
	Detail(request UserRequest) (*model.User, error)
	Delete(request UserRequest) error
	GetModel() interface{}
}
type User struct {
}
type UserRequest struct {
	pkg.CommonRequest
	ID int `form:"id" field:"id" where:"eq" default:"0"`
}

func NewUser() IUser {
	return new(User)
}
func (u *User) GetModel() interface{} {
	return model.NewUser()
}

// Add 添加
func (u *User) Add(request UserRequest) error {
	request.ReturnType = 3
	return pkg.MysqlConn.Model(u.GetModel()).Create(request.Data).Error
}

// Save 修改
func (u *User) Save(request UserRequest) error {
	request.ReturnType = 3
	return pkg.MysqlConn.Model(u.GetModel()).Updates(request.Data).Error
}
func (u *User) List(request UserRequest) (*[]model.User, error) {
	var list []model.User
	err := request.Init(pkg.MysqlConn.Model(u.GetModel()), request).Find(&list).Error
	return &list, err
}

// Detail 详情
func (u *User) Detail(request UserRequest) (*model.User, error) {
	var detail model.User
	request.ReturnType = 3
	err := request.Init(pkg.MysqlConn.Model(u.GetModel()), request).First(&detail).Error
	return &detail, err
}

// Delete 删除
func (u *User) Delete(request UserRequest) error {
	request.ReturnType = 3
	return request.Init(pkg.MysqlConn, request).Delete(u.GetModel()).Error
}
