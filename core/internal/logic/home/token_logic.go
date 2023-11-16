package home

import (
	"context"
	"fmt"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"net/http"

	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	// officialAccount *officialaccount.OfficialAccount
	r *http.Request
	w http.ResponseWriter
}

func NewTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *TokenLogic {
	// account := wechat_service.NewOfficialAccount(ctx, svcCtx.Redis, svcCtx.Config)
	return &TokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		// officialAccount: account.OfficialAccount,
		r: r,
		w: w,
	}
}

func (l *TokenLogic) Token() (resp *types.TokenResponse, err error) {
	// wechat token 处理
	// 传入request和responseWriter
	server := l.svcCtx.OfficialAccount.GetServer(l.r, l.w)
	// 设置接收消息的处理方法
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		// TODO
		// 回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	// 处理消息接收以及回复
	err = server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 发送回复的消息
	err = server.Send()
	if err != nil {
		fmt.Println(err)
		return
	}
	resp = &types.TokenResponse{}
	return
}
