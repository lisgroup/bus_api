package no_auth

import (
	"bus_api/core/define"
	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"
	"bus_api/core/service/gee"
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type GeeCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGeeCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GeeCodeLogic {
	return &GeeCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GeeCodeLogic) GeeCode(req *types.GeeRequest) (resp *types.GeeResponse, err error) {
	fmt.Println(req.Uuid)
	// 获取redis缓存的bypass状态
	bypassStatus := l.svcCtx.Redis.Get(l.ctx, define.GeeTestBypassStatusKey).String()
	var result *gee.GeetestLibResult
	gtLib := gee.NewGeetestLib(l.svcCtx.Config.GeeTestId, l.svcCtx.Config.GeeTestKey)
	digestmod := "md5"
	userID := "test"
	params := map[string]string{
		"digestmod":   digestmod,
		"user_id":     userID,
		"client_type": "web",
		"ip_address":  "127.0.0.1",
	}
	if bypassStatus == "success" {
		result = gtLib.Register(digestmod, params)
	} else {
		result = gtLib.LocalRegister()
	}
	resp = &types.GeeResponse{}
	// 解析 json
	err = json.Unmarshal([]byte(result.Data), resp)
	if err != nil {
		return
	}
	// 注意，不要更改返回的结构和值类型
	//c.Header("Content-Type", "application/json;charset=UTF-8")
	//c.String(http.StatusOK, result.Data)
	return
}
