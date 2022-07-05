package dept

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDeptLogic {
	return &DeleteDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDeptLogic) DeleteDept(req *types.DeleteDeptReq) (resp *types.DeleteDeptResp, err error) {
	userName := fmt.Sprintf("%s", l.ctx.Value("name"))
	name, err := l.svcCtx.UserModel.FindOneByName(l.ctx, userName)
	switch err {
	case nil:
		if name.RoleId != 1 {
			return nil, errors.New("非系统管理员,不能删除部门")
		}
		return &types.DeleteDeptResp{Message: "ok"}, l.svcCtx.UserDeptModel.Delete(l.ctx, req.Id)

	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", userName, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("用户登录失败,参数:%s,异常:%s", userName, err.Error())
		return nil, err
	}
}
