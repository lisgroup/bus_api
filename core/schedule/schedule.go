package schedule

import (
	"fmt"
)

// RunCrontab 运行crontab定时任务
func RunCrontab(noticeJob NoticeJob) {
	ins := NewInstance()
	_ = ins.BindTaskAndSchedule(noticeJob.Name, fmt.Sprintf("@every %ds", noticeJob.Spec), noticeJob)
}
