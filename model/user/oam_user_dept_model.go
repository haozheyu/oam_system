package user

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OamUserDeptModel = (*customOamUserDeptModel)(nil)

type (
	// OamUserDeptModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOamUserDeptModel.
	OamUserDeptModel interface {
		oamUserDeptModel
	}

	customOamUserDeptModel struct {
		*defaultOamUserDeptModel
	}
)

// NewOamUserDeptModel returns a model for the database table.
func NewOamUserDeptModel(conn sqlx.SqlConn) OamUserDeptModel {
	return &customOamUserDeptModel{
		defaultOamUserDeptModel: newOamUserDeptModel(conn),
	}
}
