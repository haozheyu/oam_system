package nav

import (
	"context"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBodyNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBodyNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBodyNameLogic {
	return &DeleteBodyNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBodyNameLogic) DeleteBodyName(req *types.DeleteBodyNameReq) (resp *types.DeleteBodyNameResp, err error) {
	// todo: add your logic here and delete this line

	return
}
