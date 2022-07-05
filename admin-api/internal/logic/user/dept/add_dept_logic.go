package dept

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

type AddDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeptLogic {
	return &AddDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddDeptLogic) AddDept(req *types.AddDeptReq) (resp *types.AddDeptResp, err error) {
	userName := fmt.Sprintf("%s", l.ctx.Value("name"))
	srcUser, err := l.svcCtx.UserModel.FindOneByName(l.ctx, userName)
	switch err {
	case nil:
		if srcUser.RoleId != 1 {
			return nil, errors.New("非系统管理员,不能添加部门")
		}
		_, err := l.svcCtx.UserDeptModel.Insert(l.ctx, &user.OamUserDept{
			Name:           req.Name,
			CreateBy:       userName,
			CreateTime:     time.Now().Unix(),
			LastUpdateBy:   userName,
			LastUpdateTime: time.Now().Unix(),
			DelFlag:        0,
		})
		return &types.AddDeptResp{Message: "ok"}, err

	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("token 解析,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, err
	}
}
