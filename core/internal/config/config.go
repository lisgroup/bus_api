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
	Salt         string
	JwtKey       string
	UserName     string
	MailPassword string
}
