package user

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/haozheyu/oam_system/admin-api/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"github.com/haozheyu/oam_system/admin-api/internal/svc"
	"github.com/haozheyu/oam_system/admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq) (resp *types.LoginResp, err error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByName(l.ctx, req.UserName)

	switch err {
	case nil:
		if userInfo.Password != config.Md5(req.Password) {
			logx.WithContext(l.ctx).Errorf("用户密码不正确,参数:%s", req.Password)
			return nil, errors.New("用户密码不正确")
		}

		now := time.Now().Unix()
		accessExpire := l.svcCtx.Config.Auth.AccessExpire
		roleDetails, _ := l.svcCtx.UserRoleModel.FindOne(l.ctx, userInfo.RoleId)
		deptDetails, _ := l.svcCtx.UserDeptModel.FindOne(l.ctx, userInfo.DeptId)
		jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id, userInfo.Name, roleDetails.Name, deptDetails.Name)
		if err != nil {
			reqStr, _ := json.Marshal(req)
			logx.WithContext(l.ctx).Errorf("生成token失败,参数:%s,异常:%s", reqStr, err.Error())
			return nil, err
		}

		reqStr, _ := json.Marshal(req)
		listStr, _ := json.Marshal(resp)
		logx.WithContext(l.ctx).Infof("登录成功,参数:%s,响应:%s", reqStr, listStr)
		return &types.LoginResp{
			Message:      "成功",
			UserName:     userInfo.Name,
			RoleName:     roleDetails.Name,
			DeptName:     deptDetails.Name,
			AccessToken:  jwtToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		}, nil
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", req.UserName, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("用户登录失败,参数:%s,异常:%s", req.UserName, err.Error())
		return nil, err
	}

	return
}

func (l *UserLoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64, name string, deptname string, rolename string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	claims["name"] = name
	claims["deptName"] = deptname
	claims["roleName"] = rolename
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
