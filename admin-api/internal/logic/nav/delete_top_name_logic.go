package nav

import (
	"context"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTopNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTopNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTopNameLogic {
	return &DeleteTopNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTopNameLogic) DeleteTopName(req *types.DeleteTopNameReq) (resp *types.DeleteTopNameResp, err error) {
	// todo: add your logic here and delete this line

	return
}
