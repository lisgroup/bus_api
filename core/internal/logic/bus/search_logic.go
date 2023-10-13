package bus

import (
	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"net/url"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchRequest) (resp *types.SearchResponse, err error) {
	// line := regexp.MustCompile(`快(\d+)`).ReplaceAllStringFunc(req.Linename, func(s string) string {
	// 	return fmt.Sprintf("快线%d号", []rune(s)[1]-'0')
	// })
	result, err := http.Get("https://szgj.2500.tv/line/search?keyword=" + req.Linename)
	if err != nil {
		return
	}
	defer result.Body.Close()
	if result.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("http get error , code=%d, status=%s", result.StatusCode, result.Status))
	}
	// 加载html对象
	doc, err := goquery.NewDocumentFromReader(result.Body)
	if err != nil {
		return
	}
	// 选择元素
	// 选择器语法就是css选择器语法，和jsoup中的类似
	resp = new(types.SearchResponse)
	doc.Find(".buslinediv .routeline").Each(func(i int, s *goquery.Selection) {
		// 获取列表数据
		buslinename := s.Find(".buslinename").Text()
		buslineto := s.Find(".buslineto").Text()
		urlStr, _ := s.Find("a").Attr("href")
		// 解析上面的url获取其中的id参数
		u, err := url.Parse(urlStr)
		if err != nil {
			return
		}
		query := u.Query()
		id := query.Get("id")
		fmt.Printf("buslinename:%s, buslineto:%s, href:%s\n", buslinename, buslineto, id)
		line := types.LineList{
			Linename: buslinename,
			Lineto:   buslineto,
			Lineid:   id,
		}
		resp.List = append(resp.List, line)
	})
	return
}
