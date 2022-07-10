package nav

import (
	"context"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBodyNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddBodyNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBodyNameLogic {
	return &AddBodyNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddBodyNameLogic) AddBodyName(req *types.AddBodyNameReq) (resp *types.AddBodyNameResp, err error) {
	// todo: add your logic here and delete this line

	return
}
