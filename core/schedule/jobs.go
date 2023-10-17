package schedule

import (
	"bus_api/core/service/push"
	"context"
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
