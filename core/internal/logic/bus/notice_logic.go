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
	"time"
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
	tx := l.svcCtx.Gorm.Model(models.Notice{}).Create(&notice)
	fmt.Println(tx)
	if tx.Error != nil {
		return nil, err
	}
	return
}

func parseTIme(receivedTime string) (minTime string, err error) {
	parsedTime, err := time.Parse(time.RFC3339, receivedTime)
	if err != nil {
		fmt.Println("时间解析错误:", err)
		return
	}

	// 将时区设置为东八区（北京时间）
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("时区加载错误:", err)
		return
	}
	parsedTime = parsedTime.In(location)

	fmt.Println("转换后的时间:", parsedTime)
	minTime = parsedTime.Format("15:04:05")
	return
}
