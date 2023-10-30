package user

import (
	"bus_api/core/xerror"
	"net/http"

	"bus_api/core/internal/logic/auth/user"
	"bus_api/core/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			httpx.ErrorCtx(r.Context(), w, xerror.NewParamsFailedError(err.Error()))
		} else {
			// httpx.OkJsonCtx(r.Context(), w, resp)
			httpx.OkJsonCtx(r.Context(), w, xerror.NewSuccessJson(resp))
		}
	}
}
