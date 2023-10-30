package member

import (
	"context"

	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogoutLogic {
	return &UserLogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogoutLogic) UserLogout() (resp *types.UserLogoutResponse, err error) {
	// 退出登录逻辑
	return
}
