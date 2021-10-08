package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID int `gorm:"primaryKey;column:id;type:int(10);default:0" json:"id"`

	AddTime    int64 `gorm:"column:add_time;type:int(10);default:0" json:"addTime"`
	UpdateTime int64 `gorm:"column:update_time;type:int(10);default:0" json:"updateTime"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (i *User) TableName() string {
	return "user"
}
func NewUser() *User {
	return new(User)
}
