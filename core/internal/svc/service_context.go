package svc

import (
	"bus_api/core/internal/config"
	"bus_api/core/internal/middleware"
	"bus_api/core/service"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Gorm   *gorm.DB
	Redis  *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	service.InitConfig(c)
	return &ServiceContext{
		Config: c,
		Gorm:   service.InitMysql(c.Mysql.Dsn, c.Mysql.MaxIdleConn, c.Mysql.MaxOpenConn, c.Mysql.MaxLifetime),
		Redis:  service.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
