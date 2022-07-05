package role

import (
	"context"
	"errors"
	"fmt"
	"github.com/haozheyu/oam_system/admin-api/model/user"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateRoleReq) (resp *types.UpdateRoleResp, err error) {
	userName := fmt.Sprintf("%s", l.ctx.Value("name"))
	name, err := l.svcCtx.UserModel.FindOneByName(l.ctx, userName)
	switch err {
	case nil:
		if name.RoleId != 1 {
			return nil, errors.New("非系统管理员,不能更新角色信息")
		}

		return &types.UpdateRoleResp{Message: "ok"}, l.svcCtx.UserRoleModel.Update(l.ctx, &user.OamUserRole{
			Id:             req.Id,
			Name:           req.Name,
			LastUpdateBy:   userName,
			LastUpdateTime: time.Now().Unix(),
			DelFlag:        0,
		})

	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", userName, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("用户登录失败,参数:%s,异常:%s", userName, err.Error())
		return nil, err
	}
}
