package dept

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

type ListDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDeptLogic {
	return &ListDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListDeptLogic) ListDept(req *types.ListDeptReq) (resp *types.ListDeptResp, err error) {
	whereBuilder := l.svcCtx.UserDeptModel.RowBuilder()
	list, err := l.svcCtx.UserDeptModel.FindPageListByPage(l.ctx, whereBuilder, 1, 50, "")
	switch err {
	case nil:
		builder := l.svcCtx.UserDeptModel.CountBuilder("*")
		count, err := l.svcCtx.UserDeptModel.FindCount(l.ctx, builder)
		if err != nil {
			l.Logger.Errorf("count: %s, err:%s", count, err)
			return nil, errors.New("部门列表获取失败")
		}
		var (
			resp_items []types.ListDeptData
			items      types.ListDeptData
		)

		for _, v := range list {

			_ = copier.Copy(&items, v)
			resp_items = append(resp_items, items)
		}
		return &types.ListDeptResp{
			Message: "ok",
			Data:    resp_items,
			Total:   count,
		}, nil
	case sqlx.ErrNotFound:
		return nil, err
	default:
		fmt.Println(err)
		return nil, errors.New("default 部门列表获取失败")
	}
}
