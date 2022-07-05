package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/haozheyu/oam_system/admin-api/internal/config"
	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReSetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReSetPasswordLogic {
	return &ReSetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReSetPasswordLogic) ReSetPassword(req *types.ReSetPasswordReq) (resp *types.ReSetPasswordResp, err error) {
	userName := fmt.Sprintf("%s", l.ctx.Value("name"))
	srcUser, err := l.svcCtx.UserModel.FindOneByName(l.ctx, userName)
	switch err {
	case nil:
		targetUser, err := l.svcCtx.UserModel.FindOneByName(l.ctx, req.UserName)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("目标用户不存在,参数:%s,异常:%s", srcUser.Name, err.Error())
			return nil, errors.New("目标用户不存在")
		}
		targetUser.Password = config.Md5(req.Password)
		targetUser.LastUpdateBy = srcUser.Name
		targetUser.LastUpdateTime = time.Now().Unix()
		return &types.ReSetPasswordResp{Message: "ok"}, l.svcCtx.UserModel.Update(l.ctx, targetUser)

	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("token 解析,参数:%s,异常:%s", srcUser.Name, err.Error())
		return nil, err
	}
}
