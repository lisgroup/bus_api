package schedule

import (
	"fmt"
	"time"
)

// RunCrontab 运行crontab定时任务
func RunCrontab() {
	name := "NoticeJob:30s"
	noticeJob := NoticeJob{
		Key:   "",
		Title: "60秒检测定时任务列表",
		Desp:  "60秒检测定时任务列表",
		Job:   Job{Name: name, CreatedAt: time.Now(), Spec: 60}}

	ins := NewInstance()
	_ = ins.BindTaskAndSchedule(noticeJob.Name, fmt.Sprintf("@every %ds", noticeJob.Spec), noticeJob)
}
