package host

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsLogic {
	return &WsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsLogic) Ws(req *types.GetWSReq) (resp *types.GetWSResp, err error) {
	fmt.Println(req.Addr, req.User, req.Port)
	builder := l.svcCtx.SshHostModel.RowBuilder()
	where := builder.Where("addr = ?", req.Addr).Where("user = ?", req.User)
	hosts, err := l.svcCtx.SshHostModel.FindOneByQuery(l.ctx, where)
	switch err {
	case nil:
		return &types.GetWSResp{req.User, req.Addr, req.Port, hosts.Password}, err
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("host不存在,参数:%s,异常:%s", req.Addr, err.Error())
		return nil, errors.New("host不存在")
	default:
		return nil, err
	}
}
