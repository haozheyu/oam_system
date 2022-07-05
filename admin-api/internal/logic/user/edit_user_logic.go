package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserLogic {
	return &EditUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditUserLogic) EditUser(req *types.EditUserReq) (resp *types.EditUserResp, err error) {
	userName := fmt.Sprintf("%s", l.ctx.Value("name"))
	srcUser, err := l.svcCtx.UserModel.FindOneByName(l.ctx, userName)
	switch err {
	case nil:
		if srcUser.RoleId != 1 {
			return nil, errors.New("非系统管理员,不能编辑用户")
		}
		targetUser, err := l.svcCtx.UserModel.FindOneByName(l.ctx, req.UserName)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("目标用户不存在,参数:%s,异常:%s", srcUser.Name, err.Error())
			return nil, errors.New("目标用户不存在")
		}

		targetUser.Name = req.UserName
		targetUser.Email = req.Email
		targetUser.Mobile = req.Mobile
		targetUser.NickName = req.NickName
		targetUser.DeptId = req.DeptId
		targetUser.RoleId = req.RoleId
		targetUser.Status = req.Status
		targetUser.Age = req.Age
		targetUser.Sex = req.Sex

		targetUser.LastUpdateBy = srcUser.Name
		targetUser.LastUpdateTime = time.Now().Unix()
		return &types.EditUserResp{Message: "ok"}, l.svcCtx.UserModel.Update(l.ctx, targetUser)

	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("token 解析,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, err
	}
}
