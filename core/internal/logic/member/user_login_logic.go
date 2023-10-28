package member

import (
	"bus_api/core/define"
	"bus_api/core/helper"
	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"
	"bus_api/core/models"
	"bus_api/core/xerror"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginRequest) (resp *types.UserResponse, err error) {
	// 登录逻辑
	// 1. 查询用户
	user := new(models.Users)
	tx := l.svcCtx.Gorm.Where("name = ?", req.Username).First(user)
	if tx.Error != nil {
		// return nil, tx.Error
		return nil, xerror.NewCodeError(xerror.RequestParamError, tx.Error.Error())
	}
	if user.Id == 0 {
		// return nil, errors.New("用户名或密码错误")
		return nil, xerror.NewCodeError(xerror.RequestParamError, "用户名或密码错误")
	}
	ok, err := helper.CheckPassword(req.Password, user.Password)
	if err != nil {
		// return nil, err
		return nil, xerror.NewCodeError(xerror.RequestParamError, err.Error())
	}
	if !ok {
		// return nil, errors.New("用户名或密码错误")
		return nil, xerror.NewCodeError(xerror.RequestParamError, "用户名或密码错误")
	}
	resp = new(types.UserResponse)
	// 2. 获取token
	resp.Token, err = helper.GenerateToken(user.Id, user.Identity, user.Username, define.TokenExpire)
	if err != nil {
		// return nil, err
		return nil, xerror.NewCodeError(xerror.RequestParamError, err.Error())
	}
	// 3. 获取刷新的token
	resp.RefreshToken, err = helper.GenerateToken(int(user.Id), user.Identity, user.Username, define.RefreshTokenExpire)
	if err != nil {
		// return nil, err
		return nil, xerror.NewCodeError(xerror.RequestParamError, err.Error())
	}
	return
}
