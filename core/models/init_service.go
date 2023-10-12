package models

import (
	"bus_api/core/define"
	"bus_api/core/internal/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// InitMysql 初始化 MySQL 配置
func InitMysql(dsn string, maxIdle, maxOpen, maxLifetime int) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 连接池设置
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetMaxOpenConns(maxOpen)
	sqlDB.SetConnMaxLifetime(time.Hour * time.Duration(maxLifetime))
	return db
}

// InitRedis 初始化 Redis 配置
func InitRedis(addr, pwd string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
}

// InitConfig 配置项设置
func InitConfig(c config.Config) {
	define.JwtKey = c.JwtKey
	define.Salt = c.Salt
	define.UserName = c.UserName
	define.MailPassword = c.MailPassword
}
