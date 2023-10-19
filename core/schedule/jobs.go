package schedule

import (
	"bus_api/core/models"
	"bus_api/core/service"
	"bus_api/core/service/push"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"time"
)

type Job struct {
	Name      string    `json:"job_name"`
	CreatedAt time.Time `json:"created_at"`
	Spec      int       `json:"spec"`
}

type NoticeJob struct {
	Key   string `json:"key"`
	Title string
	Desp  string
	Job
}

// Run 根据定时时长执行通知任务
func (c NoticeJob) Run() {
	ctx := context.Background()
	// 查询列表检测定时任务
	var notices []models.Notice
	tx := models.Gorm.Model(models.Notice{}).Find(&notices)
	if tx.Error != nil {
		logc.Error(ctx, tx.Error)
		return
	}
	serv := service.NewBusService()
	for _, notice := range notices {
		if notice.Cycle == "day" {
			// 判断是否5分钟内的任务
			if inMinute(int(notice.Hour), int(notice.Minute), 30*time.Minute) {
				// 执行查询车次任务
				resp, err := serv.RealtimeBusLine(notice.LineId)
				if err != nil {
					logc.Error(ctx, err)
					continue
				}
				// 判断map中是否有对应的
				if info, ok := resp[notice.StationId]; ok {
					realBus := ""
					for _, inTime := range info {
						realBus += inTime.BusInfo + " " + inTime.InTime
					}
					// 通知对应的用户到站了
					// 根据设置信息通知消息
					title := fmt.Sprintf("线路:%s, 站台:%s 有公交到站了", notice.LineName, notice.StationName)
					desc := fmt.Sprintf("线路:%s, 方向:%s, 站台:%s 有公交到站了，公交信息：%s", notice.LineName, notice.LineFromTo, notice.StationName, realBus)
					server := push.NewServerJ(push.ServerJParam{
						Key:   notice.JKey,
						Title: title,
						Desp:  desc,
					})
					fmt.Println(title, desc)
					err = server.Push()
					if err != nil {
						// 记录日志
						logc.Error(ctx, err.Error())
					}
				}
			}
		}
	}
	// var spec int
	// switch req.Cycle {
	// case "day":
	//
	// }
}

func inMinute(hour, min int, minute time.Duration) bool {
	currentTime := time.Now()
	targetTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), hour, min, 0, 0, currentTime.Location())

	beforeTime := targetTime.Add(-minute)
	afterTime := targetTime.Add(minute)

	if currentTime.After(beforeTime) && currentTime.Before(afterTime) {
		fmt.Printf("当前时间在的前后%d分钟范围内", minute/time.Minute)
		return true
	}
	return false
}
