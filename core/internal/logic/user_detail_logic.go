package logic

import (
	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"
	"bus_api/core/models"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {
	resp = &types.UserDetailResponse{}
	// 获取用户信息
	u := &models.Users{}
	err = l.svcCtx.Gorm.Where("identity = ?", req.Identity).First(u).Error
	if err != nil {
		return nil, err
	}
	if u.Id == 0 {
		return nil, errors.New("用户不存在")
	}
	resp.Name = u.Name
	resp.Email = u.Email
	return
}
