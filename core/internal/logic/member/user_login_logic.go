package member

import (
	"bus_api/core/define"
	"bus_api/core/helper"
	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"
	"bus_api/core/models"
	"bus_api/core/service/gee"
	"bus_api/core/xerror"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
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
	// 极验验证码
	var result *gee.GeetestLibResult
	gtLib := gee.NewGeetestLib(l.svcCtx.Config.GeeTestId, l.svcCtx.Config.GeeTestKey)
	challenge := req.GeeTestChallenge
	validate := req.GeeTestValidate
	secCode := req.GeeTestSeccode
	bypassStatus, err := l.svcCtx.Redis.Get(l.ctx, define.GeeTestBypassStatusKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		// return nil, err
		return nil, xerror.NewCodeError(xerror.RequestParamError, err.Error())
	}
	if bypassStatus == "success" {
		result = gtLib.SuccessValidate(challenge, validate, secCode)
	} else {
		// 宕机模式，直接当做验证成功
		result = gtLib.FailValidate(challenge, validate, secCode)
	}
	// 注意，不要更改返回的结构和值类型
	if result.Status != 1 {
		// 校验不通过
		return nil, xerror.NewCodeError(xerror.RequestParamError, "验证码错误")
	}

	// 登录逻辑
	// 1. 查询用户
	user := new(models.Users)
	tx := l.svcCtx.Gorm.Where("username = ?", req.Username).First(user)
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
	resp.AccessToken, err = helper.GenerateToken(user.Id, user.Identity, user.Username, define.TokenExpire, []string{user.Role})
	if err != nil {
		// return nil, err
		return nil, xerror.NewCodeError(xerror.RequestParamError, err.Error())
	}
	// 3. 获取刷新的token
	resp.RefreshToken, err = helper.GenerateToken(user.Id, user.Identity, user.Username, define.RefreshTokenExpire, []string{user.Role})
	if err != nil {
		// return nil, err
		return nil, xerror.NewCodeError(xerror.RequestParamError, err.Error())
	}
	resp.ExpiresIn = define.TokenExpire
	resp.TokenType = "Bearer"
	return
}
