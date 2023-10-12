package test

import (
	"bus_api/core/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestGorm(t *testing.T) {
	dsn := "root:@tcp(127.0.0.1:3306)/bus_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 连接池设置
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	user := &models.Users{Name: "zz", Password: "111", DeletedAt: nil}
	tx := db.Create(user)
	if tx.Error != nil {
		t.Fatalf("insert err: %v\n", err)
	}
	// 查询数据
	var res models.Users
	tx = db.Where("name = ?", "zz").First(&res)
	if tx.Error != nil {
		t.Fatalf("insert err: %v\n", err)
	}
	assert.Equal(t, user.Id, res.Id)
}
