package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	preUser, err := l.svcCtx.UserModel.FindOneByName(l.ctx, req.UserName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("用户信息获取失败：%s", err))
	}
	preUserRole, _ := l.svcCtx.UserRoleModel.FindOne(l.ctx, preUser.RoleId)
	preUserDetp, _ := l.svcCtx.UserDeptModel.FindOne(l.ctx, preUser.DeptId)
	return &types.UserInfoResp{
		Name:           preUser.Name,
		NickName:       preUser.NickName,
		Avatar:         preUser.Avatar,
		Email:          preUser.Email,
		Mobile:         preUser.Mobile,
		Status:         preUser.Status,
		DeptId:         preUser.DeptId,
		CreateBy:       preUser.CreateBy,
		CreateTime:     preUser.CreateTime,
		LastUpdateBy:   preUser.LastUpdateBy,
		LastUpdateTime: preUser.LastUpdateTime,
		RoleId:         preUser.RoleId,
		Sex:            preUser.Sex,
		Age:            preUser.Age,
		DeptName:       preUserDetp.Name,
		RoleName:       preUserRole.Name,
	}, nil
}
