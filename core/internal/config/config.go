package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql struct {
		Dsn         string
		MaxIdleConn int
		MaxOpenConn int
		MaxLifetime int
	}
	Redis struct {
		Addr     string
		Password string
		DB       int
	}
	// 微信配置
	WeChat struct {
		AppId          string `json:"appId" yaml:"AppID"`
		AppSecret      string `json:"appSecret" yaml:"AppSecret"`
		Token          string `json:"token" yaml:"Token"`
		EncodingAESKey string `json:"encodingAESKey" yaml:"EncodingAESKey"`
	} `json:"wechat" yaml:"WeChat"`
	AppUrl       string
	ApiUrl       string
	Salt         string
	JwtKey       string
	UserName     string
	MailPassword string

	// 极验配置
	GeeTestId  string
	GeeTestKey string
}
