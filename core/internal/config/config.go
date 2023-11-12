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
	AppUrl       string
	Salt         string
	JwtKey       string
	UserName     string
	MailPassword string

	// 极验配置
	GeeTestId  string
	GeeTestKey string
	// 微信配置
	Wechat struct {
		AppId          string
		AppSecret      string
		Token          string
		EncodingAESKey string
	}
}
