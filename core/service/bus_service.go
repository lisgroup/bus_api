package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type BusService struct {
}

func NewBusService() *BusService {
	return &BusService{}
}

type Line struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		NextBus   string                 `json:"nextBus"`
		NextShift string                 `json:"nextShift"`
		StandInfo map[string][]BusInTime `json:"standInfo"`
		Number    int                    `json:"number"`
	} `json:"data"`
}

type BusInTime struct {
	BusInfo string `json:"busInfo"`
	InTime  string `json:"inTime"`
}

func (b *BusService) RealtimeBusLine(lineId string) (resp map[string][]BusInTime, err error) {
	// 2. 获取接口内容
	// 根据sid查找名称
	res, err := http.Get("https://szgj.2500.tv/api/v1/busline/bus?line_guid=" + lineId)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("http get error , code=%d, status=%s", res.StatusCode, res.Status))
		return
	}
	// 读取body
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		err = errors.New(fmt.Sprintf("read body error:%s", err))
		return
	}
	// 解析json数据
	bs := Line{}
	err = json.Unmarshal(bytes, &bs)
	if err != nil {
		err = errors.New(fmt.Sprintf("json unmarshal error:%s", err))
		return
	}
	// 解析 StandInfo 数据
	if bs.Code == "0" {
		return bs.Data.StandInfo, nil
	}
	return nil, errors.New(bs.Msg)
}
