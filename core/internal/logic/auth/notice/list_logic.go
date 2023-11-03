package notice

import (
	"bus_api/core/define"
	"bus_api/core/models"
	"context"
	"github.com/jinzhu/copier"

	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.NoticeListRequest) (resp *types.NoticeListResponse, err error) {
	// 根据条件查询用户列表
	// 0. 获取分页参数
	pageSize := req.PageSize
	if pageSize == 0 {
		pageSize = define.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * pageSize
	// 1. 查询用户
	notice := new([]models.Notice)
	resp = new(types.NoticeListResponse)
	tx := l.svcCtx.Gorm
	if len(req.Keyword) > 0 {
		tx = tx.Where("line_name like ?", "%"+req.Keyword+"%").
			Or("line_from_to like ?", "%"+req.Keyword+"%").
			Or("station_name like ?", "%"+req.Keyword+"%")
	}
	// 分页查询
	var total int64
	err = tx.Model(notice).Count(&total).Error
	if err != nil {
		return
	}
	if total == 0 {
		return
	}
	err = tx.Limit(pageSize).Offset(offset).Find(&notice).Error
	if err != nil {
		return
	}
	resp.Total = total
	var list []types.Notice
	err = copier.Copy(&list, &notice)
	if err != nil {
		return
	}
	resp.List = list
	return
}
