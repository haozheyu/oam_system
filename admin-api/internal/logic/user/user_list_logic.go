package user

import (
	"context"
	"fmt"
	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.ListUserReq) (resp *types.ListUserResp, err error) {

	whereBuilder := l.svcCtx.UserModel.RowBuilder()
	list, err := l.svcCtx.UserModel.FindPageListByPage(l.ctx, whereBuilder, req.Current, req.PageSize, "")
	switch err {
	case nil:
		builder := l.svcCtx.UserModel.CountBuilder("*")
		count, err := l.svcCtx.UserModel.FindCount(l.ctx, builder)
		if err != nil {
			l.Logger.Errorf("count: %s, err:%s", count, err)
			return nil, errors.New("用户用户列表数量获取失败")
		}
		roleBuilder := l.svcCtx.UserRoleModel.RowBuilder()
		role_list, err := l.svcCtx.UserRoleModel.FindPageListByPage(l.ctx, roleBuilder, 1, 50, "")
		if err != nil {
			return nil, errors.New("角色列表获取失败")
		}
		deptBuilder := l.svcCtx.UserDeptModel.RowBuilder()
		detp_list, err := l.svcCtx.UserDeptModel.FindPageListByPage(l.ctx, deptBuilder, 1, 50, "")
		if err != nil {
			return nil, errors.New("部门列表获取失败")
		}
		var (
			resp_items []types.ListUserData
			items      types.ListUserData
		)

		for _, v := range list {
			for _, v1 := range role_list {
				if v.RoleId == v1.Id {
					items.RoleName = v1.Name
				}
			}
			for _, v1 := range detp_list {
				if v.DeptId == v1.Id {
					items.DeptName = v1.Name
				}
			}
			_ = copier.Copy(&items, v)
			resp_items = append(resp_items, items)
		}
		return &types.ListUserResp{
			Page:     req.Current,
			Data:     resp_items,
			PageSize: req.PageSize,
			Total:    count,
		}, nil
	case sqlx.ErrNotFound:
		return nil, err
	default:
		fmt.Println(err)
		return nil, errors.New("用户列表失败")
	}

}
