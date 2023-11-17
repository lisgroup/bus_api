package wechat_service

import (
	"bus_api/core/internal/config"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/openplatform"
	openConfig "github.com/silenceper/wechat/v2/openplatform/config"
	openOfficialAccount "github.com/silenceper/wechat/v2/openplatform/officialaccount"
)

// OfficialAccount 公众号操作样例
type OfficialAccount struct {
	Wechat              *wechat.Wechat
	OfficialAccount     *officialaccount.OfficialAccount
	OpenPlatform        *openplatform.OpenPlatform
	OpenOfficialAccount *openOfficialAccount.OfficialAccount
}

var wc *wechat.Wechat

// InitWechat 获取wechat实例
// 在这里已经设置了全局cache，则在具体获取公众号/小程序等操作实例之后无需再设置，设置即覆盖
func InitWechat(ctx context.Context, conn *redis.Client) *wechat.Wechat {
	if wc != nil {
		return wc
	}
	wc = wechat.NewWechat()
	redisCache := NewRedis(ctx, conn)
	wc.SetCache(redisCache)
	return wc
}

func NewOfficialAccount(ctx context.Context, conn *redis.Client, c config.Config) *OfficialAccount {
	if wc == nil {
		wc = wechat.NewWechat()
		redisCache := NewRedis(ctx, conn)
		wc.SetCache(redisCache)
	}
	redisCache := NewRedis(ctx, conn)
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

func NewOpenAccount(ctx context.Context, conn *redis.Client, c config.Config) (*openplatform.OpenPlatform, *openOfficialAccount.OfficialAccount) {
	if wc == nil {
		wc = wechat.NewWechat()
		redisCache := NewRedis(ctx, conn)
		wc.SetCache(redisCache)
	}
	redisCache := NewRedis(ctx, conn)
	cfg := &openConfig.Config{
		AppID:          c.WeChat.AppId,
		AppSecret:      c.WeChat.AppSecret,
		Token:          c.WeChat.Token,
		EncodingAESKey: c.WeChat.EncodingAESKey,
		Cache:          redisCache,
	}
	// officialAccount := wc.GetOfficialAccount(cfg)
	// 下面文档中提到的openPlatform都是这个变量
	openPlatform := wc.GetOpenPlatform(cfg)
	officialAccount := openPlatform.GetOfficialAccount(c.WeChat.AppId)
	return openPlatform, officialAccount
}
