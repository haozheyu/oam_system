package user

import (
	"github.com/haozheyu/oam_system/admin-api/internal/config"
	"net/http"

	"github.com/haozheyu/oam_system/admin-api/internal/logic/user"
	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EditUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewEditUserLogic(r.Context(), svcCtx)
		resp, err := l.EditUser(&req)
		config.Response(w, resp, err)
	}
}
