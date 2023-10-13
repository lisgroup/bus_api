package bus

import (
	"bus_api/core/xerror"
	"net/http"

	"bus_api/core/internal/logic/bus"
	"bus_api/core/internal/svc"
	"bus_api/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LineHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LineRequest
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			httpx.ErrorCtx(r.Context(), w, xerror.NewParamsFailedError(err.Error()))
			return
		}

		l := bus.NewLineLogic(r.Context(), svcCtx)
		resp, err := l.Line(&req)
		if err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			httpx.ErrorCtx(r.Context(), w, xerror.NewParamsFailedError(err.Error()))
		} else {
			// httpx.OkJsonCtx(r.Context(), w, resp)
			httpx.OkJsonCtx(r.Context(), w, xerror.NewSuccessJson(resp))
		}
	}
}
