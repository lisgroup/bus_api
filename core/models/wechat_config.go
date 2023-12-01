package models

import (
	"gorm.io/gorm"
	"time"
)

type WechatConfig struct {
	Id             int            `gorm:"primary_key;auto_increment"`
	AppName        string         `gorm:"app_name;size:100;unique:uni_app_name"`
	AppId          string         `gorm:"app_id;size:100;unique:uni_app_id"`
	AppSecret      string         `gorm:"app_secret;size:100;"`
	Token          string         `gorm:"token;size:100;"`
	EncodingAesKey string         `gorm:"encoding_aes_key;size:200;"`
	Menu           string         `gorm:"menu;size:2000;"` // 菜单
	CreatedAt      time.Time      `gorm:"created_at"`
	UpdatedAt      time.Time      `gorm:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"deleted_at"`
}

func (s *WechatConfig) TableName() string {
	return "wechat_config"
}
