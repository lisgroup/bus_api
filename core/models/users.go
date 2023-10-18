package models

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Id          int
	Identity    string         `gorm:"identity;unique:uni_identity"`
	Name        string         `gorm:"name;unique:uni_name"`
	Password    string         `gorm:"password"`
	Email       string         `gorm:"email;unique:uni_email"`
	NowVolume   int64          `gorm:"now_volume"`
	TotalVolume int64          `gorm:"total_volume"`
	CreatedAt   time.Time      `gorm:"created_at"`
	UpdatedAt   time.Time      `gorm:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"deleted_at"`
}

func (u *Users) TableName() string {
	return "users"
}
