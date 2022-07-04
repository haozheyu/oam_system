package role

import (
	"github.com/haozheyu/oam_system/admin-api/internal/config"
	"net/http"

	"github.com/haozheyu/oam_system/admin-api/internal/logic/user/role"
	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewListRoleLogic(r.Context(), svcCtx)
		resp, err := l.ListRole(&req)
		config.Response(w, resp, err)
	}
}
