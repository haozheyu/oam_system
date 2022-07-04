package user

import (
	"github.com/haozheyu/oam_system/admin-api/internal/config"
	"net/http"

	"github.com/haozheyu/oam_system/admin-api/internal/logic/user"
	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserDeleteLogic(r.Context(), svcCtx)
		resp, err := l.UserDelete(&req)
		config.Response(w, resp, err)
	}
}
