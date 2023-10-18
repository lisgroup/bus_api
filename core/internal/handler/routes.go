// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	authuser "bus_api/core/internal/handler/auth/user"
	bus "bus_api/core/internal/handler/bus"
	home "bus_api/core/internal/handler/home"
	member "bus_api/core/internal/handler/member"
	"bus_api/core/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/search",
				Handler: bus.SearchHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/line",
				Handler: bus.LineHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notice_setting",
				Handler: bus.NoticeHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/bus"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: home.HomeHandler(serverCtx),
			},
		},
		rest.WithPrefix("/"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: member.UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/mail/code/send/register",
				Handler: member.MailCodeSendRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: member.UserLoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/list",
					Handler: authuser.UserListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/detail",
					Handler: authuser.UserDetailHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/user"),
	)
}
