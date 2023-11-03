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

func (l *UpdateLogic) Update(req *types.NoticeUpdateRequest) (resp *types.NoticeUpdateResponse, err error) {
	// Server酱通知更新操作
	// 判断类型
	if req.Id <= 0 || req.Cycle != "day" && req.Cycle != "hour" && req.Cycle != "hour-n" && req.Cycle != "weekday" {
		return nil, errors.New("参数错误")
	}
	// 获取时间参数
	if len(req.StartTime) != 8 {
		req.StartTime, err = parseTIme(req.StartTime)
		if err != nil {
			return nil, err
		}
	}
	if len(req.EndTime) != 8 {
		req.EndTime, err = parseTIme(req.EndTime)
		if err != nil {
			return nil, err
		}
	}
	notice := models.Notice{}

	err = copier.Copy(&notice, &req)
	if err != nil {
		return nil, err
	}
	// 根据id更新
	tx := l.svcCtx.Gorm.Model(models.Notice{}).Where("id = ?", req.Id).Updates(&notice)
	// fmt.Println(tx)
	if tx.Error != nil {
		return nil, err
	}
	return
}
