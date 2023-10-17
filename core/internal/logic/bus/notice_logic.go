package bus

import (
	"context"

	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NoticeLogic {
	return &NoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NoticeLogic) Notice(req *types.NoticeRequest) (resp *types.NoticeResponse, err error) {
	// Server酱通知
	// 周期 { value: 'day', label: '每天' },
	//        { value: 'hour', label: '每小时' },
	//        { value: 'hour-n', label: 'N小时' },
	//        { value: 'month', label: '工作日' }

	return
}
