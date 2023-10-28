package models

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Id          int
	Identity    string         `gorm:"identity;size:36;unique:uni_identity"`
	Username    string         `gorm:"username;size:60;unique:uni_username"`
	Password    string         `gorm:"password;size:128"`
	Email       string         `gorm:"email;size:100;unique:uni_email"`
	NowVolume   int64          `gorm:"now_volume"`
	TotalVolume int64          `gorm:"total_volume"`
	CreatedAt   time.Time      `gorm:"created_at"`
	UpdatedAt   time.Time      `gorm:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"deleted_at"`
}

func (u *Users) TableName() string {
	return "users"
}
