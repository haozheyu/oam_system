package user

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ OamUserModel = (*customOamUserModel)(nil)

type (
	// OamUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOamUserModel.
	OamUserModel interface {
		oamUserModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, data *OamUser) error
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*OamUser, error)
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*OamUser, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*OamUser, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*OamUser, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*OamUser, error)
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

// export logic
func (m *defaultOamUserModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.Trans(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

// export logic
func (m *defaultOamUserModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(oamUserRows).From(m.table)
}

// export logic
func (m *defaultOamUserModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

// export logic
func (m *defaultOamUserModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultOamUserModel) DeleteSoft(ctx context.Context, data *OamUser) error {
	data.Status = 0
	data.LastUpdateTime = time.Now().Unix()
	if err := m.Update(ctx, data); err != nil {
		return errors.Wrapf(errors.New("删除数据失败"), "OamUser delete err : %+v", err)
	}
	return nil
}

func (m *defaultOamUserModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*OamUser, error) {

	query, values, err := rowBuilder.Where("status = ?", 1).ToSql()
	if err != nil {
		return nil, err
	}

	var resp OamUser
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

func (m *defaultOamUserModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error) {

	query, values, err := sumBuilder.Where("status = ?", 1).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultOamUserModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultOamUserModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*OamUser, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where("status = ?", UserStatusEnable).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*OamUser
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultOamUserModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*OamUser, error) {

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

	var resp []*OamUser
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultOamUserModel) FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64) ([]*OamUser, error) {
	query, values, err := rowBuilder.Where("status = ?", UserStatusEnable).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*OamUser
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		fmt.Println(resp)
		return resp, nil
	default:
		return nil, err
	}
}

//按照id升序分页查询数据，不支持排序
func (m *defaultOamUserModel) FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*OamUser, error) {

	if preMaxId > 0 {
		rowBuilder = rowBuilder.Where(" id > ? ", preMaxId)
	}

	query, values, err := rowBuilder.Where("status = ?", UserStatusEnable).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*OamUser
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
