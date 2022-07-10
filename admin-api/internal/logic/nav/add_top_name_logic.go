package nav

import (
	"context"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTopNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTopNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTopNameLogic {
	return &AddTopNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTopNameLogic) AddTopName(req *types.AddTopNameReq) (resp *types.AddTopNameResp, err error) {
	// todo: add your logic here and delete this line

	return
}
