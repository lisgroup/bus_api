package schedule

import (
	"testing"
	"time"
)

func TestRunCrontab(t *testing.T) {
	// noticeJob := NoticeJob{
	// 	Key:   "",
	// 	Title: "测试",
	// 	Desp:  "测试",
	// 	Job:   Job{Name: "NoticeJob", CreatedAt: time.Now(), Spec: 1800}}
	RunCrontab()
	// 睡眠100秒，等待任务执行
	time.Sleep(100 * time.Second)
}
