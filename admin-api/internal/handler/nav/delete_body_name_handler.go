package nav

import (
	"github.com/haozheyu/oam_system/admin-api/internal/config"
	"net/http"

	"github.com/haozheyu/oam_system/admin-api/internal/logic/nav"
	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteBodyNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteBodyNameReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := nav.NewDeleteBodyNameLogic(r.Context(), svcCtx)
		resp, err := l.DeleteBodyName(&req)
		config.Response(w, resp, err)
	}
}
