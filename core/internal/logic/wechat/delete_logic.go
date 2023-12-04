package wechat

import (
	"bus_api/core/models"
	"context"
	"errors"

	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.WechatDeleteRequest) (resp *types.WechatDeleteResponse, err error) {
	// 参数校验
	if req.Id <= 0 {
		return nil, errors.New("参数错误")
	}
	// 删除
	err = l.svcCtx.Gorm.Where("id = ?", req.Id).Delete(&models.WechatConfig{}).Error
	return
}
