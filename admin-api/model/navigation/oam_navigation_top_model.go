package navigation

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OamNavigationTopModel = (*customOamNavigationTopModel)(nil)

type (
	// OamNavigationTopModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOamNavigationTopModel.
	OamNavigationTopModel interface {
		oamNavigationTopModel
	}

	customOamNavigationTopModel struct {
		*defaultOamNavigationTopModel
	}
)

// NewOamNavigationTopModel returns a model for the database table.
func NewOamNavigationTopModel(conn sqlx.SqlConn) OamNavigationTopModel {
	return &customOamNavigationTopModel{
		defaultOamNavigationTopModel: newOamNavigationTopModel(conn),
	}
}
