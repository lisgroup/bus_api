package models

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Id          int
	Identity    string         `gorm:"identity;size:36;unique:uni_identity"`
	Username    string         `gorm:"username;size:60;unique:uni_username"`
	Password    string         `gorm:"password;size:128;default:''"`
	Email       string         `gorm:"email;size:100;unique:uni_email"`
	Role        string         `gorm:"role;size:100;default:'Customer'"`    // 角色
	Openid      string         `gorm:"openid;size:100;default:''"`          // 微信openid
	Unionid     string         `gorm:"unionid;size:100;default:''"`         // 微信unionid
	NowVolume   int64          `gorm:"now_volume;default:0;type:int(11)"`   // 当前使用
	TotalVolume int64          `gorm:"total_volume;default:0;type:int(11)"` // 总的
	CreatedAt   time.Time      `gorm:"created_at"`
	UpdatedAt   time.Time      `gorm:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"deleted_at"`
}

func (u *Users) TableName() string {
	return "users"
}
