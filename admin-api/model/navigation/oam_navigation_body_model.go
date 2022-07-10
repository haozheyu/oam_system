package navigation

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OamNavigationBodyModel = (*customOamNavigationBodyModel)(nil)

type (
	// OamNavigationBodyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOamNavigationBodyModel.
	OamNavigationBodyModel interface {
		oamNavigationBodyModel
	}

	customOamNavigationBodyModel struct {
		*defaultOamNavigationBodyModel
	}
)

// NewOamNavigationBodyModel returns a model for the database table.
func NewOamNavigationBodyModel(conn sqlx.SqlConn) OamNavigationBodyModel {
	return &customOamNavigationBodyModel{
		defaultOamNavigationBodyModel: newOamNavigationBodyModel(conn),
	}
}
