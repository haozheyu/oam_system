package nav

import (
	"context"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTopNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTopNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTopNameLogic {
	return &ListTopNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTopNameLogic) ListTopName(req *types.ListTopNameReq) (resp *types.ListTopNameResp, err error) {
	// todo: add your logic here and delete this line

	return
}
