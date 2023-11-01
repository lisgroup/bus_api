package schedule

import (
	"fmt"
	"time"
)

// RunCrontab 运行crontab定时任务
func RunCrontab() {
	// 1. 60秒检测定时任务列表
	name := "NoticeJob:30s"
	noticeJob := NoticeJob{
		Key:   "",
		Title: "60秒检测定时任务列表",
		Desp:  "60秒检测定时任务列表",
		Job:   Job{Name: name, CreatedAt: time.Now(), Spec: 60}}

	ins := NewInstance()
	_ = ins.BindTaskAndSchedule(noticeJob.Name, fmt.Sprintf("@every %ds", noticeJob.Spec), noticeJob)
	// 2. 轮询发送获取bypass状态请求的时间间隔10秒
	// geeTestJob := GeeTestJob{
	// 	Job{Name: "GeeTestJob:100s", CreatedAt: time.Now(), Spec: define.GeeTestCycleTime},
	// }
	// _ = ins.BindTaskAndSchedule(geeTestJob.Name, fmt.Sprintf("@every %ds", geeTestJob.Spec), geeTestJob)
}
