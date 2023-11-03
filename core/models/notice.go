package models

import (
	"gorm.io/gorm"
	"time"
)

// Notice 通知任务
type Notice struct {
	Id     int
	UserId int    `gorm:"user_id;index:idx_user_id"`
	JKey   string `gorm:"j_key;size:60;index:j_key"`
	Cycle  string `gorm:"cycle;size:20;"`
	Hour   int8   `gorm:"hour;"`
	Minute int8   `gorm:"minute;"`

	LineId      string `gorm:"line_id;size:60"`
	LineName    string `gorm:"line_name;size:60"`
	LineFromTo  string `gorm:"line_from_to;size:100"`
	StationNum  string `gorm:"station_num;size:10"`
	StationId   string `gorm:"station_id;size:60"`
	StationName string `gorm:"station_name;size:60"`

	StartTime  string `gorm:"start_time;size:20"`
	EndTime    string `gorm:"end_time;size:20"`
	NoticeTime int8   `gorm:"notice_time;size:20"`

	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at"`
}

func (s *Notice) TableName() string {
	return "notice"
}
