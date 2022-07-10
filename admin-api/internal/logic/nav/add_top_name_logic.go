package nav

import (
	"context"
	"errors"
	"fmt"
	"github.com/haozheyu/oam_system/admin-api/model/navigation"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTopNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTopNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTopNameLogic {
	return &AddTopNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTopNameLogic) AddTopName(req *types.AddTopNameReq) (resp *types.AddTopNameResp, err error) {
	userName := fmt.Sprintf("%s", l.ctx.Value("name"))
	srcUser, err := l.svcCtx.UserModel.FindOneByName(l.ctx, userName)
	switch err {
	case nil:
		if srcUser.RoleId != 1 {
			return nil, errors.New("非系统管理员,不能添加导航")
		}
		_, err := l.svcCtx.NavTopModel.Insert(l.ctx, &navigation.OamNavigationTop{
			Name:       req.Name,
			CreateTime: time.Now().Unix(),
			CreateBy:   srcUser.Name,
			Status:     1,
		})
		return &types.AddTopNameResp{Message: "ok"}, err

	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("token 解析,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, err
	}
}
