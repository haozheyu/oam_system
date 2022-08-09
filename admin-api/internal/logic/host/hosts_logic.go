package host

import (
	"context"
	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type HostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HostsLogic {
	return &HostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HostsLogic) Hosts(req *types.ListHostReq) (resp *types.ListHostResp, err error) {

	builder := l.svcCtx.SshHostModel.RowBuilder()
	countBuilder := l.svcCtx.SshHostModel.CountBuilder("*")
	count, _ := l.svcCtx.SshHostModel.FindCount(l.ctx, countBuilder)
	hosts, _ := l.svcCtx.SshHostModel.FindPageListByPage(l.ctx, builder, req.Page, req.PageSize, "")
	var (
		resp_data_item types.AddHostReq
		resp_data_list []types.AddHostReq
	)
	for _, v := range hosts {
		copier.Copy(&resp_data_item, &v)
		resp_data_list = append(resp_data_list, resp_data_item)
	}
	return &types.ListHostResp{Message: "ok", Data: resp_data_list, Total: count, Page: req.Page, PageSize: req.PageSize}, err

}
