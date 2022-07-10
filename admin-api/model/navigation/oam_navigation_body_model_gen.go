// Code generated by goctl. DO NOT EDIT!

package navigation

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	oamNavigationBodyFieldNames          = builder.RawFieldNames(&OamNavigationBody{})
	oamNavigationBodyRows                = strings.Join(oamNavigationBodyFieldNames, ",")
	oamNavigationBodyRowsExpectAutoSet   = strings.Join(stringx.Remove(oamNavigationBodyFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	oamNavigationBodyRowsWithPlaceHolder = strings.Join(stringx.Remove(oamNavigationBodyFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	oamNavigationBodyModel interface {
		Insert(ctx context.Context, data *OamNavigationBody) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*OamNavigationBody, error)
		Update(ctx context.Context, newData *OamNavigationBody) error
		Delete(ctx context.Context, id int64) error
	}

	defaultOamNavigationBodyModel struct {
		conn  sqlx.SqlConn
		table string
	}

	OamNavigationBody struct {
		Id         int64  `db:"id"`
		TopId      int64  `db:"top_id"`      // 标题ID
		LinkAddr   string `db:"link_addr"`   // 链接地址
		LinkName   string `db:"link_name"`   // 链接名称
		LinkDesc   string `db:"link_desc"`   // 链接描述
		Status     int64  `db:"status"`      // 状态
		LinkIcon   string `db:"link_icon"`   // 链接icon
		CreateTime int64  `db:"create_time"` // 创建时间
		CreateBy   string `db:"create_by"`   // 创建人
	}
)

func newOamNavigationBodyModel(conn sqlx.SqlConn) *defaultOamNavigationBodyModel {
	return &defaultOamNavigationBodyModel{
		conn:  conn,
		table: "`oam_navigation_body`",
	}
}

func (m *defaultOamNavigationBodyModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultOamNavigationBodyModel) FindOne(ctx context.Context, id int64) (*OamNavigationBody, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", oamNavigationBodyRows, m.table)
	var resp OamNavigationBody
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOamNavigationBodyModel) Insert(ctx context.Context, data *OamNavigationBody) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, oamNavigationBodyRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.TopId, data.LinkAddr, data.LinkName, data.LinkDesc, data.Status, data.LinkIcon, data.CreateBy)
	return ret, err
}

func (m *defaultOamNavigationBodyModel) Update(ctx context.Context, data *OamNavigationBody) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, oamNavigationBodyRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.TopId, data.LinkAddr, data.LinkName, data.LinkDesc, data.Status, data.LinkIcon, data.CreateBy, data.Id)
	return err
}

func (m *defaultOamNavigationBodyModel) tableName() string {
	return m.table
}
