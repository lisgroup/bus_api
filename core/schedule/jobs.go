package schedule

import (
	"bus_api/core/models"
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
	for _, notice := range notices {
		if notice.Cycle == "day" {
			// 判断是否5分钟内的任务
		}
	}
	// var spec int
	// switch req.Cycle {
	// case "day":
	//
	// }
	// 根据设置信息通知消息 "SCT204585TYnW8nBPYBPoYoGeWaG6kap6j"
	server := push.NewServerJ(push.ServerJParam{
		Key:   c.Key,
		Title: c.Title,
		Desp:  c.Desp,
	})
	err := server.Push()
	if err != nil {
		// 记录日志
		logc.Error(ctx, err.Error())
	}
}

func inFiveMinute(hour, min int) bool {
	currentTime := time.Now()
	targetTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), hour, min, 0, 0, currentTime.Location())

	beforeTime := targetTime.Add(-5 * time.Minute)
	afterTime := targetTime.Add(5 * time.Minute)

	if currentTime.After(beforeTime) && currentTime.Before(afterTime) {
		fmt.Println("当前时间在的前后五分钟范围内")
		return true
	}
	return false
}
