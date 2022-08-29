package ssh

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SshHostsModel = (*customSshHostsModel)(nil)

type (
	// SshHostsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSshHostsModel.
	SshHostsModel interface {
		sshHostsModel
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		DeleteStatus(ctx context.Context, addr string) error
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*SshHosts, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]SshHosts, error)
	}

	customSshHostsModel struct {
		*defaultSshHostsModel
	}
)

// NewSshHostsModel returns a model for the database table.
func NewSshHostsModel(conn sqlx.SqlConn) SshHostsModel {
	return &customSshHostsModel{
		defaultSshHostsModel: newSshHostsModel(conn),
	}
}

// export logic
func (m *defaultSshHostsModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(sshHostsRows).From(m.table)
}

// export logic
func (m *defaultSshHostsModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT("+field+")").From(m.table).Where("idDel = ?", 0)
}

func (m *defaultSshHostsModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultSshHostsModel) DeleteStatus(ctx context.Context, addr string) error {
	one, err := m.FindOneByAddr(ctx, addr)
	switch err {
	case nil:
		one.IdDel = "1"
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

func (m *defaultSshHostsModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*SshHosts, error) {
	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp SshHosts
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

func (m *defaultSshHostsModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]SshHosts, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := rowBuilder.Where("idDel = ?", 0).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var (
		resp []SshHosts
	)
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	fmt.Println(err, resp, "----------")
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
