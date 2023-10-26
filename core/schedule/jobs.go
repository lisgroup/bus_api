package schedule

import (
	"bus_api/core/define"
	"bus_api/core/models"
	"bus_api/core/service"
	"bus_api/core/service/push"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
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
	if tx.Error != nil && tx.Error != redis.Nil {
		logc.Error(ctx, tx.Error)
		return
	}
	serv := service.NewBusService()
	for _, notice := range notices {
		if notice.Cycle == "day" || notice.Cycle == "one" {
			// 是否已经完成通知了
			noticeKey := time.Now().Format("20060102") + "_notice_time:" + strconv.Itoa(notice.Id)
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
					// err = models.Gorm.Model(&models.Notice{}).Where("id = ?", notice.Id).UpdateColumn("end_at", current).Error
					// if err != nil {
					// 	logc.Error(ctx, err)
					// }
				}
				continue
			}

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
				url := define.AppUrl + "/#/line?lineID=" + notice.LineId + "&to=" + notice.LineFromTo + "&lineName=" + notice.LineName
				desc := fmt.Sprintf(`      [线路]:%s
      [方向]:%s
      [站台]:%s 到站了
      [公交信息]:%s
      [查看详情]:[%s](%s)`, notice.LineName, notice.LineFromTo, notice.StationName, realBus, notice.LineName+notice.LineFromTo, url)
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
					err = models.Redis.Set(ctx, noticeKey, noticeTime+1, time.Hour*24).Err()
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
