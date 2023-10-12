package logic

import (
	"bus_api/core/define"
	"bus_api/core/models"
	"context"
	"github.com/jinzhu/copier"

	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserRequest) (resp *types.UserListResponse, err error) {
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
	user := new([]models.Users)
	resp = new(types.UserListResponse)
	tx := l.svcCtx.Gorm
	if len(req.Name) > 0 {
		tx = tx.Where("name like ?", "%"+req.Name+"%")
	}
	// 分页查询
	var total int64
	err = tx.Model(user).Count(&total).Error
	if err != nil {
		return
	}
	if total == 0 {
		return
	}
	err = tx.Limit(pageSize).Offset(offset).Find(&user).Error
	if err != nil {
		return
	}
	resp.Total = total
	var list []types.User
	err = copier.Copy(&list, &user)
	if err != nil {
		return
	}
	resp.List = list
	return
}
