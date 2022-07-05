package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/haozheyu/oam_system/admin-api/internal/config"
	"github.com/haozheyu/oam_system/admin-api/model/user"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req *types.AddUserReq) (resp *types.AddUserResp, err error) {
	userName := fmt.Sprintf("%s", l.ctx.Value("name"))
	srcUser, err := l.svcCtx.UserModel.FindOneByName(l.ctx, userName)
	switch err {
	case nil:
		if srcUser.RoleId != 1 {
			return nil, errors.New("非系统管理员,不能添加用户")
		}
		_, err := l.svcCtx.UserModel.Insert(l.ctx, &user.OamUser{
			Name:           req.Name,
			NickName:       req.NickName,
			Avatar:         req.Avatar,
			Password:       config.Md5(req.Password),
			Email:          req.Email,
			Mobile:         req.Mobile,
			Status:         1,
			DeptId:         req.DeptId,
			CreateBy:       userName,
			CreateTime:     time.Now().Unix(),
			LastUpdateBy:   userName,
			LastUpdateTime: time.Now().Unix(),
			DelFlag:        0,
			RoleId:         req.RoleId,
			Sex:            req.Sex,
			Age:            req.Age,
		})
		return &types.AddUserResp{Message: "ok"}, err

	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("token 解析,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, err
	}
}
