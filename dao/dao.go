package dao

import (
	"crawler/conf"
	"crawler/utils"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Dao struct {
	db    *sql.DB
	idGen *utils.IdGenService
}

func NewDao(c *conf.Config, idGen *utils.IdGenService) *Dao {
	dao := &Dao{}
	dao.db = getSqliteDB(c.Sqlite.Dsn)
	dao.idGen = idGen

	err := dao.db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("sqlite ping ok")

	cronsRow, err := dao.CreateCronsTable()
	if err != nil {
		panic(err)
	}
	fmt.Println("CreateCronsTable ok", cronsRow)

	cronsLogRow, err := dao.CreateCronsLogTable()
	if err != nil {
		panic(err)
	}
	fmt.Println("CreateCronsLogTable ok", cronsLogRow)
	return dao
}

func (d *Dao) Close() {
	d.db.Close()
}

func getSqliteDB(dsn string) (db *sql.DB) {
	var err error
	db, err = sql.Open("sqlite3", dsn)
	if err != nil {
		panic(err)
	}

	return
}

//创建表
const (
	cronsTable = `
	CREATE TABLE IF NOT EXISTS crons (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name varchar(64) not null,
		status INTEGER not null,
		expr varchar(64) not null,
		command text null,
		created_at datetime null,
		updated_at datetime null
	);
`
	cronsLogTable = `
	CREATE TABLE IF NOT EXISTS crons_log (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		cron_name varchar(64) not null,
		exec_time datetime not null,
		http_code INTEGER not null,
		http_context_size INTEGER not null,
		http_context text not null,
		created_at datetime null
	);
	`
)

func (d *Dao) CreateCronsTable() (id int64, err error) {
	res, err := d.db.Exec(cronsTable)
	if err != nil {
		return
	}
	if res == nil {
		return
	}
	id, err = res.RowsAffected()
	return
}

func (d *Dao) CreateCronsLogTable() (id int64, err error) {
	res, err := d.db.Exec(cronsLogTable)
	if err != nil {
		return
	}
	if res == nil {
		return
	}
	id, err = res.RowsAffected()
	return
}
