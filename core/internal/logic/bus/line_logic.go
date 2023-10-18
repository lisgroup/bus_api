package bus

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"

	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LineLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type Line struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		NextBus   string `json:"nextBus"`
		NextShift string `json:"nextShift"`
		StandInfo map[string][]struct {
			BusInfo string `json:"busInfo"`
			InTime  string `json:"inTime"`
		} `json:"standInfo"`
		Number int `json:"number"`
	} `json:"data"`
}

func NewLineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LineLogic {
	return &LineLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LineLogic) Line(req *types.LineRequest) (resp *types.LineResponse, err error) {
	// 公交详情--所有站台列表
	// 1. 获取html内容
	// map 记录站台id和公交详情列表的索引
	var m = map[string]int{}
	// 发送htt get请求 https://szgj.2500.tv/line/bus?id=0000000000LINELINEINFO18122459783711
	result, err := http.Get("https://szgj.2500.tv/line/bus?id=" + req.Lineid)
	if err != nil {
		return
	}
	defer result.Body.Close()
	if result.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("http get error , code=%d, status=%s", result.StatusCode, result.Status))
		return
	}
	// 加载html对象
	doc, err := goquery.NewDocumentFromReader(result.Body)
	if err != nil {
		return
	}
	// 获取列表数据
	resp = new(types.LineResponse)
	// 选择器语法就是css选择器语法，和jsoup中的类似
	doc.Find(".busline .stationinfo").Each(func(i int, s *goquery.Selection) {
		// 获取文本
		num := s.Find(".stationnum").Text()
		name := s.Find(".stationname").Text()
		// businfo := s.Find(".businfo").Text()
		// 获取属性值
		sid, _ := s.Find(".stationdetail").Attr("data-sid")
		// 存入map方便查找
		m[sid] = len(resp.List)
		list := types.LineInfo{
			Stationnum:    num,
			Stationname:   name,
			Stationid:     sid,
			Stationdetail: []string{}, // 需要后面api接口更新
		}
		resp.List = append(resp.List, list)
	})

	// 2. 获取接口内容
	// 根据sid查找名称
	res, err := http.Get("https://szgj.2500.tv/api/v1/busline/bus?line_guid=" + req.Lineid)
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
		for sid, infos := range bs.Data.StandInfo {
			for _, v := range infos {
				// 根据 sid 获取站点名称
				idx := m[sid]
				// 更新站点进站公交列表
				resp.List[idx].Stationdetail = append(resp.List[idx].Stationdetail, v.BusInfo+" "+v.InTime)
			}
		}
	}
	return
}
