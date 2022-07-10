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

type ListBodyNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListBodyNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBodyNameLogic {
	return &ListBodyNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListBodyNameLogic) ListBodyName(req *types.ListBodyNameReq) (resp *types.ListBodyNameResp, err error) {
	userName := fmt.Sprintf("%s", l.ctx.Value("name"))
	srcUser, err := l.svcCtx.UserModel.FindOneByName(l.ctx, userName)
	switch err {
	case nil:
		if srcUser.RoleId != 1 {
			return nil, errors.New("非系统管理员,不能添加导航")
		}
		builder := l.svcCtx.NavBodyModel.RowBuilder()
		page, _ := l.svcCtx.NavBodyModel.FindPageListByPage(l.ctx, builder, req.Page, req.PageSize, "")
		countBuilder := l.svcCtx.NavBodyModel.CountBuilder("*")
		count, _ := l.svcCtx.NavBodyModel.FindCount(l.ctx, countBuilder)
		navTopLsit, _ := l.svcCtx.NavTopModel.FindPageListByPage(l.ctx, l.svcCtx.NavTopModel.RowBuilder())
		var (
			resp_data_item types.ListBodyNameData
			resp_data_list []types.ListBodyNameData
		)
		for _, v := range page {
			for _, v1 := range navTopLsit {
				if v.TopId == v1.Id {
					copier.Copy(resp_data_item, v)
					resp_data_item.TopName = v1.Name
					resp_data_list = append(resp_data_list, resp_data_item)
				}
			}
		}
		return &types.ListBodyNameResp{Message: "ok", Data: resp_data_list, Total: count}, err

	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("token 解析,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, err
	}
}
