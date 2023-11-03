package notice

import (
	"bus_api/core/models"
	"context"
	"errors"
	"github.com/jinzhu/copier"

	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.NoticeDetailRequest) (resp *types.NoticeDetailResponse, err error) {
	// 获取详情数据
	notice := models.Notice{}
	// 获取用户id
	userId := l.ctx.Value("id").(int)
	tx := l.svcCtx.Gorm.
		Where("id = ?", req.Id).
		Where("user_id = ?", userId).First(&notice)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("没有找到该通知")
	}
	resp = &types.NoticeDetailResponse{}
	err = copier.Copy(resp, &notice)
	if err != nil {
		return nil, err
	}
	return
}
