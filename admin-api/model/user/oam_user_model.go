package user

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OamUserModel = (*customOamUserModel)(nil)

type (
	// OamUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOamUserModel.
	OamUserModel interface {
		oamUserModel
	}

	customOamUserModel struct {
		*defaultOamUserModel
	}
)

// NewOamUserModel returns a model for the database table.
func NewOamUserModel(conn sqlx.SqlConn) OamUserModel {
	return &customOamUserModel{
		defaultOamUserModel: newOamUserModel(conn),
	}
}
