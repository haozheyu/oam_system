package role

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRoleLogic {
	return &ListRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListRoleLogic) ListRole(req *types.ListRoleReq) (resp *types.ListRoleResp, err error) {
	whereBuilder := l.svcCtx.UserRoleModel.RowBuilder()
	list, err := l.svcCtx.UserRoleModel.FindPageListByPage(l.ctx, whereBuilder, 1, 50, "")
	switch err {
	case nil:
		builder := l.svcCtx.UserRoleModel.CountBuilder("*")
		count, err := l.svcCtx.UserRoleModel.FindCount(l.ctx, builder)
		if err != nil {
			l.Logger.Errorf("count: %s, err:%s", count, err)
			return nil, errors.New("角色列表获取失败")
		}
		var (
			resp_items []types.ListRoleData
			items      types.ListRoleData
		)

		for _, v := range list {

			_ = copier.Copy(&items, v)
			resp_items = append(resp_items, items)
		}
		return &types.ListRoleResp{
			Message: "ok",
			Data:    resp_items,
			Total:   count,
		}, nil
	case sqlx.ErrNotFound:
		return nil, err
	default:
		fmt.Println(err)
		return nil, errors.New("default 角色列表获取失败")
	}
}
