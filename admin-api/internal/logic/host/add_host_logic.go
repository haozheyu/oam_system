package host

import (
	"context"
	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"
	"github.com/haozheyu/oam_system/admin-api/model/ssh"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddHostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHostLogic {
	return &AddHostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddHostLogic) AddHost(req *types.AddHostReq) (resp *types.AddHostResp, err error) {
	_, err = l.svcCtx.SshHostModel.Insert(l.ctx, &ssh.SshHosts{
		Addr:     req.Addr,
		User:     req.User,
		Port:     req.Port,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddHostResp{Message: "ok"}, err
}
