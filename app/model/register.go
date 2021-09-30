package model

import (
	"gorm.io/gorm"
	"time"
)

type RegisterCenter struct {
	ID int `gorm:"primaryKey;column:id;type:int(10);default:0" json:"id"`

	Name  string `gorm:"index;column:name;type:varchar(50);default:" json:"name" form:"name"`
	Title string `gorm:"column:title;type:varchar(50);default:" json:"title" form:"title"`
	Port  string `gorm:"column:port;type:varchar(20);default:" json:"port" form:"port"`
	Host  string `gorm:"column:host;type:varchar(50);default:" json:"host" form:"host"`

	AddTime    int64 `gorm:"column:add_time;type:int(10);default:0" json:"addTime"`
	UpdateTime int64 `gorm:"column:update_time;type:int(10);default:0" json:"updateTime"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (r *RegisterCenter) TableName() string {
	return "register_center"
}
func NewRegisterCenter() *RegisterCenter {
	return new(RegisterCenter)
}
