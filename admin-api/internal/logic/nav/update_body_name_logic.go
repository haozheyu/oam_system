package nav

import (
	"context"
	"errors"
	"fmt"
	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBodyNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBodyNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBodyNameLogic {
	return &UpdateBodyNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBodyNameLogic) UpdateBodyName(req *types.UpdateBodyNameReq) (resp *types.UpdateBodyNameResp, err error) {
	userName := fmt.Sprintf("%s", l.ctx.Value("name"))
	srcUser, err := l.svcCtx.UserModel.FindOneByName(l.ctx, userName)
	switch err {
	case nil:
		if srcUser.RoleId != 1 {
			return nil, errors.New("非系统管理员,不能添加导航")
		}
		link_name, err := l.svcCtx.NavBodyModel.FindOneByQuery(l.ctx, l.svcCtx.NavBodyModel.RowBuilder().Where("link_name = ?", req.LinkName))
		if err != nil {
			return nil, errors.New("查询标题内容失败")
		}
		link_name.LinkAddr = req.LinkAddr
		link_name.LinkName = req.LinkName
		link_name.LinkIcon = req.LinkIcon
		link_name.LinkDesc = req.LinkDesc
		err = l.svcCtx.NavBodyModel.Update(l.ctx, link_name)
		if err != nil {
			return nil, errors.New("更新标题内容失败")
		}
		return &types.UpdateBodyNameResp{Message: "ok"}, err

	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("token 解析,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, err
	}
}
