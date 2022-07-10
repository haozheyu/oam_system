package nav

import (
	"context"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListBodyNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListBodyNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBodyNameLogic {
	return &ListBodyNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListBodyNameLogic) ListBodyName(req *types.ListBodyNameReq) (resp *types.ListBodyNameResp, err error) {
	// todo: add your logic here and delete this line

	return
}
