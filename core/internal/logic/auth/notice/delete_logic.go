package notice

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

func (l *DeleteLogic) Delete(req *types.NoticeDeleteRequest) (resp *types.NoticeDeleteResponse, err error) {
	// 参数校验
	if req.Id <= 0 {
		return nil, errors.New("参数错误")
	}
	// 删除通知
	tx := l.svcCtx.Gorm.Where("id = ?", req.Id).Delete(&models.Notice{})
	if tx.Error != nil {
		return nil, tx.Error
	}
	return
}
