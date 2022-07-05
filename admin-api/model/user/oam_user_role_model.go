package user

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OamUserRoleModel = (*customOamUserRoleModel)(nil)

type (
	// OamUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOamUserRoleModel.
	OamUserRoleModel interface {
		oamUserRoleModel

		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*OamUserRole, error)
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

func (m *defaultOamUserRoleModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := countBuilder.Where("status = ?", 1).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultOamUserRoleModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*OamUserRole, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := rowBuilder.Where("status = ?", UserStatusEnable).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*OamUserRole
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultOamUserRoleModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultOamUserRoleModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(oamUserRoleRows).From(m.table)
}
