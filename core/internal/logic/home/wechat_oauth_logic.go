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
	// 发起网页授权--公众号
	oauth := l.svcCtx.OfficialAccount.GetOauth()
	// 重定向到微信oauth授权登录
	redirectURI := l.svcCtx.Config.ApiUrl + "/wechat_oauth_callback"
	// scope 应用授权作用域，snsapi_base （不弹出授权页面直接跳转只能获取用户openid）
	// snsapi_userinfo （弹出授权页面可通过openid拿到昵称、性别、所在地。并且， 即使在未关注的情况下，只要用户授权，也能获取其信息 ）
	// err = oauth.Redirect(l.w, l.r, redirectURI, "snsapi_userinfo", "", l.svcCtx.Config.WeChat.AppId)
	err = oauth.Redirect(l.w, l.r, redirectURI, "snsapi_userinfo", "")
	if err != nil {
		return nil, err
	}
	return
}
