package svc

import (
	"github.com/haozheyu/oam_system/admin-api/internal/config"
	"github.com/haozheyu/oam_system/admin-api/model/user"
)

type ServiceContext struct {
	Config            config.Config
	RDB               *redis.Client
	UserModel         user.OamUserModel
	UserDeptModel     user.OamUserDeptModel
	UserRoleModel     user.OamUserRoleModel
	UserRoleDeptModel user.OamUserRoleDeptModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
