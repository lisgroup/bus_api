package models

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Id          int
	Identity    string
	Name        string
	Password    string
	Email       string
	NowVolume   int64          `gorm:"now_volume"`
	TotalVolume int64          `gorm:"total_volume"`
	CreatedAt   time.Time      `gorm:"created"`
	UpdatedAt   time.Time      `gorm:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"deleted"`
}

func (u *Users) TableName() string {
	return "users"
}
