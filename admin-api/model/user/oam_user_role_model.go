package user

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OamUserRoleModel = (*customOamUserRoleModel)(nil)

type (
	// OamUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOamUserRoleModel.
	OamUserRoleModel interface {
		oamUserRoleModel
	}

	customOamUserRoleModel struct {
		*defaultOamUserRoleModel
	}
)

// NewOamUserRoleModel returns a model for the database table.
func NewOamUserRoleModel(conn sqlx.SqlConn) OamUserRoleModel {
	return &customOamUserRoleModel{
		defaultOamUserRoleModel: newOamUserRoleModel(conn),
	}
}
