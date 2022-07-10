package nav

import (
	"context"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBodyNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBodyNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBodyNameLogic {
	return &UpdateBodyNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBodyNameLogic) UpdateBodyName(req *types.UpdateBodyNameReq) (resp *types.UpdateBodyNameResp, err error) {
	// todo: add your logic here and delete this line

	return
}
