package models

import (
	"gorm.io/gorm"
	"time"
)

type ServerKey struct {
	Id        int
	UserId    int            `gorm:"user_id;index:idx_user_id"`
	JKey      string         `gorm:"j_key;size:60;unique:j_key"`
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at"`
}

func (s *ServerKey) TableName() string {
	return "server_key"
}
