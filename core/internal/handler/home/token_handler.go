package home

import (
	"net/http"

	"bus_api/core/internal/logic/home"
	"bus_api/core/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := home.NewTokenLogic(r.Context(), svcCtx, w, r)
		_, err := l.Token()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
	}
}
