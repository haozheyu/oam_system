package nav

import (
	"context"
	"errors"
	"fmt"
	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"
	"github.com/haozheyu/oam_system/admin-api/model/navigation"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBodyNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddBodyNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBodyNameLogic {
	return &AddBodyNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddBodyNameLogic) AddBodyName(req *types.AddBodyNameReq) (resp *types.AddBodyNameResp, err error) {
	userName := fmt.Sprintf("%s", l.ctx.Value("name"))
	srcUser, err := l.svcCtx.UserModel.FindOneByName(l.ctx, userName)
	switch err {
	case nil:
		if srcUser.RoleId != 1 {
			return nil, errors.New("非系统管理员,不能添加导航")
		}
		_, err := l.svcCtx.NavBodyModel.Insert(l.ctx, &navigation.OamNavigationBody{
			TopId:      req.TopNameId,
			LinkAddr:   req.LinkAddr,
			LinkName:   req.LinkName,
			LinkDesc:   req.LinkDesc,
			Status:     1,
			LinkIcon:   req.LinkIcon,
			CreateTime: time.Now().Unix(),
			CreateBy:   srcUser.Name,
		})
		return &types.AddBodyNameResp{Message: "ok"}, err

	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("token 解析,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, err
	}
}
