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

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.WechatUpdateRequest) (resp *types.WechatUpdateResponse, err error) {
	// 参数校验 AppName，AppId 不能为空，且除req.Id外不能重复
	if req.Id <= 0 || len(req.AppName) == 0 || len(req.AppId) == 0 {
		err = errors.New("app_name or app_id is empty")
		return
	}
	// 1. 查询AppName，AppId是否存在
	wechat := models.WechatConfig{}
	err = l.svcCtx.Gorm.Model(models.WechatConfig{}).Where("id != ? and (app_name = ? or app_id = ?)", req.Id, req.AppName, req.AppId).First(&wechat).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if wechat.Id > 0 {
		err = errors.New("app_name or app_id is exist")
		return
	}
	// 2. 更新微信配置
	err = copier.Copy(&wechat, &req)
	if err != nil {
		return
	}
	err = l.svcCtx.Gorm.Model(models.WechatConfig{}).Where("id = ?", req.Id).Updates(&req).Error
	if err != nil {
		return
	}
	resp = &types.WechatUpdateResponse{
		Id: wechat.Id,
	}
	return
}
