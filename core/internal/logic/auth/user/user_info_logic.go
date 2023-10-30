package user

import (
	"bus_api/core/models"
	"context"

	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// 获取用户信息
	// 1. 获取用户信息
	userId := l.ctx.Value("id").(int)
	// identity := l.ctx.Value("identity").(string)
	name := l.ctx.Value("name").(string)
	// 2. 获取用户角色
	// 查询用户表信息
	var user models.Users
	err = l.svcCtx.Gorm.Model(&models.Users{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		return
	}
	resp = &types.UserInfoResponse{
		Roles: []string{user.Role},
		Name:  name,
		// Identity: identity,
		Avatar: "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
	}
	return
}
