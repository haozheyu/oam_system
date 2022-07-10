package navigation

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OamNavigationTopModel = (*customOamNavigationTopModel)(nil)

type (
	// OamNavigationTopModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOamNavigationTopModel.
	OamNavigationTopModel interface {
		oamNavigationTopModel
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		DeleteStatus(ctx context.Context, id int64) error
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*OamNavigationTop, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]OamNavigationTop, error)
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

// export logic
func (m *defaultOamNavigationTopModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(oamNavigationTopRows).From(m.table)
}

// export logic
func (m *defaultOamNavigationTopModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultOamNavigationTopModel) DeleteStatus(ctx context.Context, id int64) error {
	one, err := m.FindOne(ctx, id)
	switch err {
	case nil:
		one.Status = 0
		if err := m.Update(ctx, one); err != nil {
			return errors.Wrapf(errors.New("删除数据失败"), "OamUser delete err : %+v", err)
		}
		return nil
	case sqlx.ErrNotFound:
		return errors.New("未查询到记录")
	default:
		return errors.New("查询信息有误")
	}
}

func (m *defaultOamNavigationTopModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*OamNavigationTop, error) {
	query, values, err := rowBuilder.Where("status = ?", 1).ToSql()
	if err != nil {
		return nil, err
	}

	var resp OamNavigationTop
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

func (m *defaultOamNavigationTopModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {
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

func (m *defaultOamNavigationTopModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]OamNavigationTop, error) {
	query, values, err := rowBuilder.Where("status = ?", 1).ToSql()
	if err != nil {
		return nil, err
	}
	var resp []OamNavigationTop
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
