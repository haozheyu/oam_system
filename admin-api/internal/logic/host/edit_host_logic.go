package host

import (
	"context"
	"github.com/haozheyu/oam_system/admin-api/model/ssh"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditHostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditHostLogic {
	return &EditHostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditHostLogic) EditHost(req *types.EditHostReq) (resp *types.EditHostResp, err error) {
	hosts, err := l.svcCtx.SshHostModel.FindOneByAddr(l.ctx, req.Addr)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.SshHostModel.Update(l.ctx, &ssh.SshHosts{
		Id:       hosts.Id,
		Addr:     hosts.Addr,
		User:     req.User,
		Port:     req.Port,
		Password: req.Password,
		IdDel:    req.IsDel,
	})
	if err != nil {
		return nil, err
	}
	return &types.EditHostResp{Message: "ok"}, err
}
