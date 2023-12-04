package wechat

import (
	"bus_api/core/models"
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.WechatAddRequest) (resp *types.WechatUpdateResponse, err error) {
	// 参数校验 AppName，AppId 不能为空，且不能重复
	if len(req.AppName) == 0 || len(req.AppId) == 0 {
		err = errors.New("AppName or AppId is empty")
		return
	}
	// 1. 查询AppName，AppId是否存在
	wechat := models.WechatConfig{}
	err = l.svcCtx.Gorm.Model(models.WechatConfig{}).Where("app_name = ? or app_id = ?", req.AppName, req.AppId).First(&wechat).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if wechat.Id > 0 {
		err = errors.New("app_name or app_id is exist")
		return
	}
	// 2. 创建微信配置
	err = copier.Copy(&wechat, &req)
	if err != nil {
		return
	}
	err = l.svcCtx.Gorm.Model(models.WechatConfig{}).Create(&wechat).Error
	if err != nil {
		return
	}
	resp = &types.WechatUpdateResponse{
		Id: wechat.Id,
	}
	return
}
