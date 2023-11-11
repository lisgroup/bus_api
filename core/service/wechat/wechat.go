package wechat

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/silenceper/wechat/v2"
)

// InitWechat 获取wechat实例
// 在这里已经设置了全局cache，则在具体获取公众号/小程序等操作实例之后无需再设置，设置即覆盖
func InitWechat(ctx context.Context, conn *redis.Client) *wechat.Wechat {
	wc := wechat.NewWechat()
	redisCache := NewRedis(ctx, conn)
	wc.SetCache(redisCache)
	return wc
}
