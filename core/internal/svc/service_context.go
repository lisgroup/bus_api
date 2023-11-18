package svc

import (
	"bus_api/core/internal/config"
	"bus_api/core/internal/middleware"
	"bus_api/core/service"
	"bus_api/core/service/wechat_service"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/openplatform"
	openOfficialAccount "github.com/silenceper/wechat/v2/openplatform/officialaccount"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config              config.Config
	Gorm                *gorm.DB
	Redis               *redis.Client
	Auth                rest.Middleware
	Wechat              *wechat.Wechat
	OfficialAccount     *officialaccount.OfficialAccount
	OpenPlatform        *openplatform.OpenPlatform
	OpenOfficialAccount *openOfficialAccount.OfficialAccount
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := context.Background()
	client := service.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB)
	account := wechat_service.NewOfficialAccount(ctx, client, c)
	openPlatform, openAccount := wechat_service.NewOpenAccount(ctx, client, c)
	service.InitConfig(c)
	return &ServiceContext{
		Config:              c,
		Gorm:                service.InitMysql(c.Mysql.Dsn, c.Mysql.MaxIdleConn, c.Mysql.MaxOpenConn, c.Mysql.MaxLifetime),
		Redis:               client,
		Auth:                middleware.NewAuthMiddleware().Handle,
		Wechat:              account.Wechat,
		OfficialAccount:     account.OfficialAccount,
		OpenPlatform:        openPlatform,
		OpenOfficialAccount: openAccount,
	}
}
