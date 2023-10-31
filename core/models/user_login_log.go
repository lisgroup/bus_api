package models

import "time"

type UserLoginLog struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"column:user_id"`
	IP         string    `gorm:"column:ip"`
	LoginTime  int       `gorm:"column:login_time"`
	Address    string    `gorm:"column:address"`
	Device     string    `gorm:"column:device"`
	DeviceType string    `gorm:"column:device_type"`
	Browser    string    `gorm:"column:browser"`
	Platform   string    `gorm:"column:platform"`
	Language   string    `gorm:"column:language"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (UserLoginLog) TableName() string {
	return "user_login_logs"
}
