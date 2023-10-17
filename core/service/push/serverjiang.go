package push

import (
	"bus_api/core/service/utils"
)

func NewServerJ(param ServerJParam) *ServerJ {
	return &ServerJ{
		ServerJParam: param,
	}
}

type ServerJParam struct {
	Key   string
	Title string
	Desp  string
}

type ServerJ struct {
	ServerJParam
	request *utils.Request
}

func (p *ServerJ) Push() error {
	var url = "https://sctapi.ftqq.com/" + p.Key + ".send"
	p.request = utils.NewUtilsRequest(utils.RequestParam{
		Url: url,
		Params: map[string]string{
			"title": p.Title,
			"desp":  p.Desp,
		},
	})

	p.request.Get()
	return nil
}
