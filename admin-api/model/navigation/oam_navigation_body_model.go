package navigation

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OamNavigationBodyModel = (*customOamNavigationBodyModel)(nil)

type (
	// OamNavigationBodyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOamNavigationBodyModel.
	OamNavigationBodyModel interface {
		oamNavigationBodyModel
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		DeleteStatus(ctx context.Context, id int64) error
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*OamNavigationBody, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]OamNavigationBody, error)
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

// export logic
func (m *defaultOamNavigationBodyModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(oamNavigationBodyRows).From(m.table)
}

// export logic
func (m *defaultOamNavigationBodyModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT("+field+")").From(m.table).Where("status = ?", 1)
}

func (m *defaultOamNavigationBodyModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := countBuilder.ToSql()
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

func (m *defaultOamNavigationBodyModel) DeleteStatus(ctx context.Context, id int64) error {
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

func (m *defaultOamNavigationBodyModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*OamNavigationBody, error) {
	query, values, err := rowBuilder.Where("status = ?", 1).ToSql()
	if err != nil {
		return nil, err
	}

	var resp OamNavigationBody
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

func (m *defaultOamNavigationBodyModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]OamNavigationBody, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := rowBuilder.Where("status = ?", 1).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []OamNavigationBody
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
