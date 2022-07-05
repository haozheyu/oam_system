package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/haozheyu/oam_system/admin-api/model/user"
	"time"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UpdateUserReq) (resp *types.UpdateUserResp, err error) {
	userName := fmt.Sprintf("%s", l.ctx.Value("name"))
	preUser, err := l.svcCtx.UserModel.FindOneByName(l.ctx, userName)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", preUser.Name, err.Error())
		return nil, errors.New("用户不存在")
	}
	return &types.UpdateUserResp{Message: "成功"}, l.svcCtx.UserModel.Update(l.ctx, &user.OamUser{
		Id:             preUser.Id,
		Name:           preUser.Name,
		NickName:       req.NickName,
		Avatar:         preUser.Avatar,
		Password:       preUser.Password,
		Email:          req.Email,
		Mobile:         req.Mobile,
		Status:         req.Status,
		DeptId:         req.DeptId,
		CreateBy:       userName,
		CreateTime:     time.Now().Unix(),
		LastUpdateBy:   userName,
		LastUpdateTime: time.Now().Unix(),
		DelFlag:        preUser.DelFlag,
		RoleId:         req.RoleId,
		Sex:            req.Sex,
		Age:            req.Age,
	})
}
