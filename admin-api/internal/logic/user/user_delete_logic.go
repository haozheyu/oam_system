package user

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteLogic) UserDelete(req *types.DeleteUserReq) (resp *types.DeleteUserResp, err error) {
	name, err := l.svcCtx.UserModel.FindOneByName(l.ctx, req.UserName)
	switch err {
	case nil:
		if name.RoleId != 1 {
			return nil, errors.New("非系统管理员,不能删除用户")
		}
		return &types.DeleteUserResp{Message: "ok"}, l.svcCtx.UserModel.Delete(l.ctx, name.Id)

	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", req.UserName, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("用户登录失败,参数:%s,异常:%s", req.UserName, err.Error())
		return nil, err
	}
}
