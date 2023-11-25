package home

import (
	"bus_api/core/models"
	"context"
	"fmt"

	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WechatOauthCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWechatOauthCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WechatOauthCallbackLogic {
	return &WechatOauthCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WechatOauthCallbackLogic) WechatOauthCallback(req *types.OauthCallbackRequest) (resp *types.OauthCallbackResponse, err error) {
	// 页面将跳转至 redirect_uri/?code=CODE&state=STATE
	// 通过code换取网页授权access_token
	// oauth := l.svcCtx.OpenOfficialAccount.PlatformOauth()
	oa := l.svcCtx.OfficialAccount
	accessToken, err := oa.GetOauth().GetUserAccessToken(req.Code)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 代第三方公众号实现网页授权
	// officialAccount := l.svcCtx.OpenOfficialAccount
	// openPlatform := l.svcCtx.OpenPlatform
	// componentAccessToken, err := openPlatform.GetComponentAccessToken()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// accessToken, err := officialAccount.PlatformOauth().GetUserAccessToken(req.Code, l.svcCtx.Config.WeChat.AppId, componentAccessToken)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(accessToken)
	// 通过accessToken获取用户信息请参考微信公众号的业务
	resp = &types.OauthCallbackResponse{
		AccessToken: accessToken.AccessToken,
	}
	// 判断 users 表是否存在这个openid，不存在的插入数据，存在的返回用户信息
	var user models.Users
	err = l.svcCtx.Gorm.Model(&models.Users{}).Where("openid = ?", accessToken.OpenID).First(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if user.Id == 0 {
		// 不存在的插入数据
		user.Openid = accessToken.OpenID
		user.Username = accessToken.OpenID[0:10]
		user.Role = "Customer"
		user.Unionid = accessToken.UnionID
		err = l.svcCtx.Gorm.Create(&user).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return
}
