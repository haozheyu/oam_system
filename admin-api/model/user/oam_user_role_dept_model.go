package user

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OamUserRoleDeptModel = (*customOamUserRoleDeptModel)(nil)

type (
	// OamUserRoleDeptModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOamUserRoleDeptModel.
	OamUserRoleDeptModel interface {
		oamUserRoleDeptModel
	}

	customOamUserRoleDeptModel struct {
		*defaultOamUserRoleDeptModel
	}
)

// NewOamUserRoleDeptModel returns a model for the database table.
func NewOamUserRoleDeptModel(conn sqlx.SqlConn) OamUserRoleDeptModel {
	return &customOamUserRoleDeptModel{
		defaultOamUserRoleDeptModel: newOamUserRoleDeptModel(conn),
	}
}
