package service

import (
	"bus_api/core/define"
	"bus_api/core/internal/config"
	"bus_api/core/models"
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	Gorm  *gorm.DB
	Redis *redis.Client
)

// InitMysql 初始化 MySQL 配置
func InitMysql(dsn string, maxIdle, maxOpen, maxLifetime int) *gorm.DB {
	var err error
	Gorm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 连接池设置
	sqlDB, _ := Gorm.DB()
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetMaxOpenConns(maxOpen)
	sqlDB.SetConnMaxLifetime(time.Hour * time.Duration(maxLifetime))
	// 自动创建数据库
	autoTables := []interface{}{
		&models.Users{},
		&models.ServerKey{},
		&models.Notice{},
		&models.UserLoginLog{},
		&models.WechatConfig{},
	}
	if err := Gorm.AutoMigrate(autoTables...); err != nil {
		panic(err)
	}
	return Gorm
}

// InitRedis 初始化 Redis 配置
func InitRedis(addr, pwd string, db int) *redis.Client {
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		panic("redis init failed")
	}

	return Redis
}

// InitConfig 配置项设置
func InitConfig(c config.Config) {
	define.AppUrl = c.AppUrl
	define.ApiUrl = c.ApiUrl
	define.JwtKey = c.JwtKey
	define.Salt = c.Salt
	define.UserName = c.UserName
	define.MailPassword = c.MailPassword
	define.GeeTestId = c.GeeTestId
	define.GeeTestKey = c.GeeTestKey
}
