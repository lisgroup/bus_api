package home

import (
	"context"
	"net/http"

	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WechatOauthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
	w      http.ResponseWriter
}

func NewWechatOauthLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *WechatOauthLogic {
	return &WechatOauthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
		w:      w,
	}
}

func (l *WechatOauthLogic) WechatOauth() (resp *types.OauthResponse, err error) {
	// 发起网页授权
	oauth := l.svcCtx.OpenOfficialAccount.PlatformOauth()
	// 重定向到微信oauth授权登录
	redirectURI := "http://www.abc.com"
	oauth.Redirect(l.w, l.r, redirectURI, "snsapi_userinfo", "", l.svcCtx.Config.WeChat.AppId)
	return
}
