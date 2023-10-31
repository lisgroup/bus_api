package user

import (
	"bus_api/core/models"
	"context"
	"fmt"
	"time"

	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogLogic {
	return &UserLoginLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogLogic) UserLoginLog(req *types.LoginLogRequest) (resp *types.LoginLogResponse, err error) {
	// 根据输入的长度获取对应的登录日志
	section := int(req.Section)
	start := time.Now().AddDate(0, 0, -(section - 1))
	end := time.Now()
	// 1. 取对应日期的日志
	// $datas = DB::table('login_log')->where('login_time', '>=', $start)->where('login_time', '<=', $end)->select('id', 'ip', 'login_time')->get();
	// var data []models.UserLoginLog
	// l.svcCtx.Gorm.Model(&models.UserLoginLog{}).Where("login_time >= ? AND login_time <= ?", start, end).Select("id, ip, login_time").Find(&data)
	// sum := len(data)
	// 2. 统计每天的登录次数
	type LoginCount struct {
		LoginDate  time.Time `gorm:"column:login_date"`
		LoginCount int       `gorm:"column:login_count"`
	}
	var result []LoginCount
	// start := time.Date(2019, 10, 1, 0, 0, 0, 0, time.UTC)
	// end := time.Date(2019, 10, 7, 23, 59, 59, 0, time.UTC)
	ss := start.Format("2006-01-02")
	fmt.Println(ss, end.Format("2006-01-02"))
	l.svcCtx.Gorm.Model(&models.UserLoginLog{}).
		Select("DATE(FROM_UNIXTIME(login_time)) AS login_date, COUNT(*) AS login_count").
		Where("DATE(FROM_UNIXTIME(login_time)) BETWEEN ? AND ?", start.Format("2006-01-02"), end.Format("2006-01-02")).
		Group("login_date").
		Find(&result)
	// 初始化返回值
	resp = &types.LoginLogResponse{}
	var m = map[string]int{}
	// 遍历输出结果
	for _, entry := range result {
		m[entry.LoginDate.Format("2006-01-02")] = entry.LoginCount
		// fmt.Printf("Date: %s, Login Count: %d\n", entry.LoginDate.Format("2006-01-02"), entry.LoginCount)
		// resp.Date = append(resp.Date, entry.LoginDate.Format("2006-01-02"))
		// resp.SuccessSlide = append(resp.SuccessSlide, entry.LoginCount)
		resp.Total++
	}
	// 3. 补全没有登录的日期
	for i := 0; i < section; i++ {
		date := start.AddDate(0, 0, i).Format("2006-01-02")
		resp.Date = append(resp.Date, date)
		resp.SuccessSlide = append(resp.SuccessSlide, m[date])
	}
	return
}
