package wechat_service

import (
	"bus_api/core/internal/config"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

// OfficialAccount 公众号操作样例
type OfficialAccount struct {
	Wechat          *wechat.Wechat
	OfficialAccount *officialaccount.OfficialAccount
}

// InitWechat 获取wechat实例
// 在这里已经设置了全局cache，则在具体获取公众号/小程序等操作实例之后无需再设置，设置即覆盖
func InitWechat(ctx context.Context, conn *redis.Client) *wechat.Wechat {
	wc := wechat.NewWechat()
	redisCache := NewRedis(ctx, conn)
	wc.SetCache(redisCache)
	return wc
}

func NewOfficialAccount(ctx context.Context, conn *redis.Client, c config.Config) *OfficialAccount {
	wc := wechat.NewWechat()
	redisCache := NewRedis(ctx, conn)
	wc.SetCache(redisCache)
	cfg := &offConfig.Config{
		AppID:          c.WeChat.AppId,
		AppSecret:      c.WeChat.AppSecret,
		Token:          c.WeChat.Token,
		EncodingAESKey: c.WeChat.EncodingAESKey,
		Cache:          redisCache,
	}
	officialAccount := wc.GetOfficialAccount(cfg)
	return &OfficialAccount{
		Wechat:          wc,
		OfficialAccount: officialAccount,
	}
}
