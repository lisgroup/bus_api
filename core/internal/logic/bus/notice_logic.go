package bus

import (
	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"
	"bus_api/core/models"
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
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
	// 判断类型
	if req.Cycle != "day" && req.Cycle != "hour" && req.Cycle != "hour-n" && req.Cycle != "weekday" {
		return nil, errors.New("")
	}
	notice := models.Notice{}

	err = copier.Copy(&notice, &req)
	if err != nil {
		return nil, err
	}
	tx := l.svcCtx.Gorm.Model(models.Notice{}).Create(&notice)
	fmt.Println(tx)
	if tx.Error != nil {
		return nil, err
	}
	return
}
