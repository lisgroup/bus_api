package schedule

import (
	"bus_api/core/models"
	"bus_api/core/service"
	"bus_api/core/service/push"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"strconv"
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
	current := time.Now().Format("15:04:05")
	var notices []models.Notice
	tx := models.Gorm.Model(&models.Notice{}).Where("start_time <= ?", current).Where("end_time >= ?", current).Find(&notices)
	if tx.Error != nil {
		logc.Error(ctx, tx.Error)
		return
	}
	serv := service.NewBusService()
	for _, notice := range notices {
		if notice.Cycle == "day" || notice.Cycle == "one" {
			// 执行查询车次任务
			resp, err := serv.RealtimeBusLine(notice.LineId)
			if err != nil {
				logc.Error(ctx, err)
				continue
			}
			noticeKey := time.Now().Format("20060102") + "_notice_time:" + strconv.Itoa(notice.Id)
			// 判断map中是否有对应的
			if info, ok := resp[notice.StationId]; ok {
				// 判断通知次数是否大于限制了
				noticeTime, err := models.Redis.Get(ctx, noticeKey).Int()
				if err != nil {
					logc.Error(ctx, err)
				}
				if noticeTime >= int(notice.NoticeTime) {
					logc.Info(ctx, "通知次数上限。。。")
					if notice.Cycle == "one" {
						// 移除计划
						err = models.Gorm.Delete(&models.Notice{}, notice.Id).Error
						if err != nil {
							logc.Error(ctx, err)
						}
					} else {
						// 修改结束时间为当前时间，防止继续请求接口
						err = models.Gorm.Model(&models.Notice{}).Where("id = ?", notice.Id).UpdateColumn("end_at", current).Error
						if err != nil {
							logc.Error(ctx, err)
						}
					}
					continue
				}
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
				} else {
					// 通知次数++
					err = models.Redis.Set(ctx, noticeKey, noticeTime+1, 86400).Err()
					if err != nil {
						logc.Error(ctx, err)
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

func inMinute(hour, min int, minuteBefore, minuteAfter time.Duration) bool {
	currentTime := time.Now()
	targetTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), hour, min, 0, 0, currentTime.Location())

	beforeTime := targetTime.Add(-minuteBefore)
	afterTime := targetTime.Add(minuteAfter)

	if currentTime.After(beforeTime) && currentTime.Before(afterTime) {
		fmt.Printf("当前时间在的前后%d--%d分钟范围内", minuteBefore/time.Minute, minuteAfter/time.Minute)
		return true
	}
	return false
}
