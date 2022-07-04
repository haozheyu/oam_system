// Code generated by goctl. DO NOT EDIT!

package user

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
	oamUserFieldNames          = builder.RawFieldNames(&OamUser{})
	oamUserRows                = strings.Join(oamUserFieldNames, ",")
	oamUserRowsExpectAutoSet   = strings.Join(stringx.Remove(oamUserFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	oamUserRowsWithPlaceHolder = strings.Join(stringx.Remove(oamUserFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	oamUserModel interface {
		Insert(ctx context.Context, data *OamUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*OamUser, error)
		FindOneByEmail(ctx context.Context, email sql.NullString) (*OamUser, error)
		FindOneByName(ctx context.Context, name string) (*OamUser, error)
		Update(ctx context.Context, newData *OamUser) error
		Delete(ctx context.Context, id int64) error
	}

	defaultOamUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	OamUser struct {
		Id             int64          `db:"id"`               // 编号
		Name           string         `db:"name"`             // 用户名
		NickName       sql.NullString `db:"nick_name"`        // 昵称
		Avatar         sql.NullString `db:"avatar"`           // 头像
		Password       sql.NullString `db:"password"`         // 密码
		Email          sql.NullString `db:"email"`            // 邮箱
		Mobile         sql.NullString `db:"mobile"`           // 手机号
		Status         sql.NullInt64  `db:"status"`           // 状态  0：禁用   1：正常
		DeptId         sql.NullInt64  `db:"dept_id"`          // 部门ID
		CreateBy       sql.NullString `db:"create_by"`        // 创建人
		CreateTime     sql.NullInt64  `db:"create_time"`      // 创建时间
		LastUpdateBy   sql.NullString `db:"last_update_by"`   // 更新人
		LastUpdateTime sql.NullInt64  `db:"last_update_time"` // 更新时间
		DelFlag        int64          `db:"del_flag"`         // 是否删除  -1：已删除  0：正常
		RoleId         sql.NullInt64  `db:"role_id"`          // 岗位Id
		Sex            int64          `db:"sex"`              // 0男,1女
		Age            int64          `db:"age"`              // 默认0
	}
)

func newOamUserModel(conn sqlx.SqlConn) *defaultOamUserModel {
	return &defaultOamUserModel{
		conn:  conn,
		table: "`oam_user`",
	}
}

func (m *defaultOamUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultOamUserModel) FindOne(ctx context.Context, id int64) (*OamUser, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", oamUserRows, m.table)
	var resp OamUser
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

func (m *defaultOamUserModel) FindOneByEmail(ctx context.Context, email sql.NullString) (*OamUser, error) {
	var resp OamUser
	query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", oamUserRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, email)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOamUserModel) FindOneByName(ctx context.Context, name string) (*OamUser, error) {
	var resp OamUser
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", oamUserRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOamUserModel) Insert(ctx context.Context, data *OamUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, oamUserRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.NickName, data.Avatar, data.Password, data.Email, data.Mobile, data.Status, data.DeptId, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag, data.RoleId, data.Sex, data.Age)
	return ret, err
}

func (m *defaultOamUserModel) Update(ctx context.Context, newData *OamUser) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, oamUserRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Name, newData.NickName, newData.Avatar, newData.Password, newData.Email, newData.Mobile, newData.Status, newData.DeptId, newData.CreateBy, newData.LastUpdateBy, newData.LastUpdateTime, newData.DelFlag, newData.RoleId, newData.Sex, newData.Age, newData.Id)
	return err
}

func (m *defaultOamUserModel) tableName() string {
	return m.table
}
