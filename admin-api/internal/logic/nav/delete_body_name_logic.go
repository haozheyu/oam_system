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

type DeleteBodyNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBodyNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBodyNameLogic {
	return &DeleteBodyNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBodyNameLogic) DeleteBodyName(req *types.DeleteBodyNameReq) (resp *types.DeleteBodyNameResp, err error) {
	userName := fmt.Sprintf("%s", l.ctx.Value("name"))
	srcUser, err := l.svcCtx.UserModel.FindOneByName(l.ctx, userName)
	switch err {
	case nil:
		if srcUser.RoleId != 1 {
			return nil, errors.New("非系统管理员,不能添加导航")
		}

		link_name := l.svcCtx.NavBodyModel.RowBuilder().Where("link_name = ?", req.LinkName)
		query, err := l.svcCtx.NavBodyModel.FindOneByQuery(l.ctx, link_name)
		if err != nil {
			return nil, errors.New("删除失败")
		}
		err = l.svcCtx.NavBodyModel.DeleteStatus(l.ctx, query.Id)
		if err != nil {
			return nil, errors.New("删除失败")
		}
		return &types.DeleteBodyNameResp{Message: "ok"}, err

	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("token 解析,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, err
	}
}
