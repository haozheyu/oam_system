package nav

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTopNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTopNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTopNameLogic {
	return &ListTopNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTopNameLogic) ListTopName(req *types.ListTopNameReq) (resp *types.ListTopNameResp, err error) {
	userName := fmt.Sprintf("%s", l.ctx.Value("name"))
	srcUser, err := l.svcCtx.UserModel.FindOneByName(l.ctx, userName)
	switch err {
	case nil:
		if srcUser.RoleId != 1 {
			return nil, errors.New("非系统管理员,不能添加导航")
		}
		builder := l.svcCtx.NavTopModel.RowBuilder()
		page, _ := l.svcCtx.NavTopModel.FindPageListByPage(l.ctx, builder)
		countBuilder := l.svcCtx.NavBodyModel.CountBuilder("*")
		count, _ := l.svcCtx.NavTopModel.FindCount(l.ctx, countBuilder)
		var (
			resp_data_item types.ListTopNameData
			resp_data_list []types.ListTopNameData
		)

		for _, v := range page {
			copier.Copy(resp_data_item, v)
			resp_data_list = append(resp_data_list, resp_data_item)
		}
		return &types.ListTopNameResp{Message: "ok", Data: resp_data_list, Total: count}, err

	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("token 解析,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, err
	}
}
