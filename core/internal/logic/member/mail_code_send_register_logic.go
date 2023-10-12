package member

import (
	"bus_api/core/define"
	"bus_api/core/helper"
	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"
	"bus_api/core/models"
	"bus_api/core/xerror"
	"context"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"github.com/zeromicro/go-zero/core/logx"
	"net/smtp"
	"time"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.SendEMailRequest) (resp *types.SendEMailUserResponse, err error) {
	// 发送邮件逻辑
	// 0. 参数校验
	if req.Email == "" {
		// return nil, errors.New("参数错误")
		return nil, xerror.NewCodeError(xerror.RequestParamError, "参数错误")
	}
	// 1. 判断邮箱是否注册
	user := new(models.Users)
	l.svcCtx.Gorm.Where("email = ?", req.Email).First(user)
	if user.Id != 0 {
		// return nil, errors.New("该邮箱已注册")
		return nil, xerror.NewCodeError(xerror.RequestParamError, "该邮箱已注册")
	}
	// 2. 限制发送次数==每天最多发送5次，每分钟最多发送1次
	dayKey := "email_register:" + time.Now().Format("20060102_") + req.Email
	result, _ := l.svcCtx.Redis.Get(l.ctx, dayKey).Int()
	if result > 5 {
		// return nil, errors.New("今日发送次数已达上限")
		return nil, xerror.NewCodeError(xerror.RequestParamError, "今日发送次数已达上限")
	}
	l.svcCtx.Redis.Set(l.ctx, dayKey, result+1, time.Hour*24)
	minuteKey := "email_register_minute:" + req.Email
	duration, err := l.svcCtx.Redis.TTL(l.ctx, minuteKey).Result()
	if err != nil {
		// return nil, err
		return nil, xerror.NewCodeError(xerror.RequestParamError, err.Error())
	}
	if duration > 0 {
		// return nil, errors.New("发送过于频繁")
		return nil, xerror.NewCodeError(xerror.RequestParamError, "发送过于频繁")
	}
	l.svcCtx.Redis.Set(l.ctx, minuteKey, result, time.Minute*1)
	// 3. 生成验证码并发送邮件
	code := helper.GenerateCode(define.CodeLength)
	// 4. 将验证码存入redis
	l.svcCtx.Redis.Set(l.ctx, "email_register:"+req.Email, code, time.Minute*5)
	go func() {
		err := l.SendEMailCode(req.Email, code)
		if err != nil {
			logx.Error(err)
		}
	}()
	return
}

func (l *MailCodeSendRegisterLogic) SendEMailCode(mail, code string) error {
	e := email.NewEmail()
	e.From = "Get <" + define.UserName + ">"
	e.To = []string{mail}
	e.Subject = "测试验证码发送"
	e.HTML = []byte("你的验证码为：<h2>" + code + "</h2>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", define.UserName, define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	return nil
}
