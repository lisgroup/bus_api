package models

import (
	"gorm.io/gorm"
	"time"
)

// Notice 通知任务
type Notice struct {
	Id     int
	UserId int    `gorm:"user_id;index:idx_user_id"`
	JKey   string `gorm:"j_key;unique:j_key"`
	Cycle  string `gorm:"cycle"`
	Hour   int8   `gorm:"hour;"`
	Minute int8   `gorm:"minute;"`

	LineId      string `gorm:"line_id"`
	LineName    string `gorm:"line_name"`
	LineFromTo  string `gorm:"line_from_to"`
	StationNum  string `gorm:"station_num"`
	StationId   string `gorm:"station_id"`
	StationName string `gorm:"station_name"`

	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at"`
}

func (s *Notice) TableName() string {
	return "notice"
}
