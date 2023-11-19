package home

import (
	"net/http"

	"bus_api/core/internal/logic/home"
	"bus_api/core/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func WechatOauthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := home.NewWechatOauthLogic(r.Context(), svcCtx, w, r)
		resp, err := l.WechatOauth()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
