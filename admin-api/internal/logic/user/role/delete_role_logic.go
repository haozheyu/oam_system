package role

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRoleLogic) DeleteRole(req *types.DeleteRoleReq) (resp *types.DeleteRoleResp, err error) {
	userName := fmt.Sprintf("%s", l.ctx.Value("name"))
	name, err := l.svcCtx.UserModel.FindOneByName(l.ctx, userName)
	switch err {
	case nil:
		if name.RoleId != 1 {
			return nil, errors.New("非系统管理员,不能删除角色")
		}
		return &types.DeleteRoleResp{Message: "ok"}, l.svcCtx.UserRoleModel.Delete(l.ctx, req.Id)

	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", userName, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("用户登录失败,参数:%s,异常:%s", userName, err.Error())
		return nil, err
	}
}
