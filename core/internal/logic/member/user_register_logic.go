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

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserResponse, err error) {
	// 注册逻辑
	// 0. 参数校验
	if req.Username == "" || req.Password == "" || req.Email == "" {
		// return nil, errors.New("参数错误")
		return nil, xerror.NewCodeError(xerror.RequestParamError, "参数错误")
	}
	// 判断 code 是否和 redis 一致
	code, err := l.svcCtx.Redis.Get(l.ctx, "email_register:"+req.Email).Result()
	if err != nil {
		// return nil, errors.New("验证码已过期")
		return nil, xerror.NewCodeError(xerror.RequestParamError, "验证码已过期")
	}
	if code != req.Code {
		// return nil, errors.New("验证码错误")
		return nil, xerror.NewCodeError(xerror.RequestParamError, "验证码错误")
	}
	// 1. 查询用户是否存在
	user := new(models.Users)
	l.svcCtx.Gorm.Where("username = ?", req.Username).First(user)
	if user.Id != 0 {
		// return nil, errors.New("用户已存在")
		return nil, xerror.NewCodeError(xerror.RequestParamError, "用户已存在")
	}
	// 2. 创建用户
	user.Username = req.Username
	user.Password, err = helper.GeneratePassword(req.Password)
	// fmt.Println(user.Password, err)
	if err != nil {
		// return nil, err
		return nil, xerror.NewCodeError(xerror.RequestParamError, "密码生成错误")
	}
	user.Email = req.Email
	user.Identity = helper.UUID()
	tx := l.svcCtx.Gorm.Create(user)
	// fmt.Println(tx)
	if tx.Error != nil {
		// fmt.Println(tx.Error)
		return nil, tx.Error
	}
	resp = new(types.UserResponse)
	// 3. 获取token
	resp.Token, err = helper.GenerateToken(user.Id, user.Identity, user.Username, define.TokenExpire)
	if err != nil {
		// return nil, err
		return nil, xerror.NewCodeError(xerror.RequestParamError, "token生成错误")
	}
	// 4. 获取刷新的token
	resp.RefreshToken, err = helper.GenerateToken(user.Id, user.Identity, user.Username, define.RefreshTokenExpire)
	return
}
