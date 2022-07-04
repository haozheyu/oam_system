package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/haozheyu/oam_system/admin-api/internal/config"
	"github.com/haozheyu/oam_system/admin-api/model/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	conn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:            c,
		RDB:               user.InitRedis(c),
		UserModel:         user.NewOamUserModel(conn),
		UserDeptModel:     user.NewOamUserDeptModel(conn),
		UserRoleModel:     user.NewOamUserRoleModel(conn),
		UserRoleDeptModel: user.NewOamUserRoleDeptModel(conn),
	}
}
