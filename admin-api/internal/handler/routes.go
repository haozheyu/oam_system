// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	host "github.com/haozheyu/oam_system/admin-api/internal/handler/host"
	nav "github.com/haozheyu/oam_system/admin-api/internal/handler/nav"
	user "github.com/haozheyu/oam_system/admin-api/internal/handler/user"
	userdept "github.com/haozheyu/oam_system/admin-api/internal/handler/user/dept"
	userrole "github.com/haozheyu/oam_system/admin-api/internal/handler/user/role"
	"github.com/haozheyu/oam_system/admin-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/userInfo",
				Handler: user.UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: user.UserAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/list",
				Handler: user.UserListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: user.UserUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: user.UserDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/reSetPassword",
				Handler: user.ReSetPasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/UpdateUserStatus",
				Handler: user.EditUserHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/user/login",
				Handler: user.UserLoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: userdept.AddDeptHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/list",
				Handler: userdept.ListDeptHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: userdept.UpdateDeptHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: userdept.DeleteDeptHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/user/dept"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: userrole.AddRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/list",
				Handler: userrole.ListRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: userrole.UpdateRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: userrole.DeleteRoleHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/user/role"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/list_top",
				Handler: nav.ListTopNameHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/list_body",
				Handler: nav.ListBodyNameHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/add_top",
				Handler: nav.AddTopNameHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/add_body",
				Handler: nav.AddBodyNameHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: nav.UpdateBodyNameHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete_top",
				Handler: nav.DeleteTopNameHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete_body",
				Handler: nav.DeleteBodyNameHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/get_top_body",
				Handler: nav.GetTopBodyNameHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/nav"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ws",
				Handler: host.WsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/hosts",
				Handler: host.HostsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/addHost",
				Handler: host.AddHostHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/editHost",
				Handler: host.EditHostHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/host"),
	)
}
