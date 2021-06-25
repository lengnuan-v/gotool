// +----------------------------------------------------------------------
// | mysql
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2021年06月25日
// +----------------------------------------------------------------------

package gotool

import (
	"github.com/gohouse/gorose"
)

// Dsn "账号:密码@tcp(IP:端口)/数据库?charset=utf8"
// Driver mysql/sqlite/oracle/mssql/postgres
// Prefix 前缀
type DB struct {
	Dsn    string // dsn
	Driver string
	Prefix string // 前缀
}

func (d *DB) GetDb() (*gorose.Connection, error) {
	var config = &gorose.DbConfigSingle{
		Driver:          d.Driver,  // 驱动: mysql/sqlite/oracle/mssql/postgres
		EnableQueryLog:  true,      // 是否开启sql日志
		SetMaxOpenConns: 0,         // (连接池)最大打开的连接数，默认值为0表示不限制
		SetMaxIdleConns: 0,         // (连接池)闲置的连接数
		Prefix:          d.Prefix,  // 表前缀
		Dsn:             d.Dsn,     // 数据库链接
	}
	return gorose.Open(config)
}

// 新增数据
func (d *DB) Insert(tanleName string, data interface{}) (int64, error) {
	var err error
	var connection *gorose.Connection
	if connection, err = d.GetDb(); err != nil {
		return 0, err
	}
	defer connection.Close()
	return connection.NewSession().Table(tanleName).Data(data).InsertGetId()
}

// 更新数据
func (d *DB) Update(tanleName string, set map[string]interface{}, cond interface{}) (int64, error) {
	var err error
	var connection *gorose.Connection
	if connection, err = d.GetDb(); err != nil {
		return 0, err
	}
	defer connection.Close()
	db := connection.NewSession()
	return db.Table(tanleName).Data(set).Where(cond).Update()
}

// 删除数据
func (d *DB) Delete(tanleName string, cond interface{}) (int64, error) {
	var err error
	var connection *gorose.Connection
	if connection, err = d.GetDb(); err != nil {
		return 0, err
	}
	defer connection.Close()
	return connection.NewSession().Table(tanleName).Where(cond).Delete()
}

// 获取条数
func (d *DB) Count(tanleName string, cond interface{}) (int64, error) {
	var err error
	var connection *gorose.Connection
	if connection, err = d.GetDb(); err != nil {
		return 0, err
	}
	defer connection.Close()
	return connection.NewSession().Table(tanleName).Where(cond).Count()
}

// 查询一条数据
func (d *DB) SelectRow(tanleName string, fields []string, cond interface{}, order string) (map[string]interface{}, error) {
	var err error
	var connection *gorose.Connection
	if connection, err = d.GetDb(); err != nil {
		return nil, err
	}
	defer connection.Close()
	return connection.NewSession().Table(tanleName).Fields(Implode(",", fields)).Where(cond).OrderBy(order).First()
}

// 查询多条数据
func (d *DB) SelectRows(tanleName string, fields []string, cond interface{}, order string) ([]map[string]interface{}, error) {
	var err error
	var connection *gorose.Connection
	if connection, err = d.GetDb(); err != nil {
		return nil, err
	}
	defer connection.Close()
	return connection.NewSession().Table(tanleName).Fields(Implode(",", fields)).Where(cond).OrderBy(order).Get()
}

// 分页查询数据
func (d *DB) SelectAll(tanleName string, fields []string, cond interface{}, order string, limit, offset int) ([]map[string]interface{}, error) {
	var err error
	var connection *gorose.Connection
	if connection, err = d.GetDb(); err != nil {
		return nil, err
	}
	defer connection.Close()
	return connection.NewSession().Table(tanleName).Fields(Implode(",", fields)).Where(cond).OrderBy(order).Limit(limit).Offset(offset).Get()
}