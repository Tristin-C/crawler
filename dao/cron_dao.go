package dao

import (
	"context"
	"crawler/model"
	"crawler/utils"
	"encoding/json"
)

func (d *Dao) BatchCron(ctx context.Context, cronList model.CronsList) (ids []int64, err error) {
	if len(cronList) <= 0 {
		return
	}

	sql := "insert into crons(id, name, status, expr, command, created_at, updated_at) values(?,?,?,?,?,?,?)"

	stmt, err := d.db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	tx, err := d.db.Begin()
	if err != nil {
		tx.Rollback()
		return
	}
	for _, item := range cronList {
		id := d.idGen.GetId()
		_, err = stmt.Exec(id, item.Name, item.Status, item.Expr, utils.ToJson(item.Command), item.CreatedAt, item.UpdatedAt)
		if err != nil {
			tx.Rollback()
			return
		}
		ids = append(ids, id)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return
	}

	return
}

func (d *Dao) CronInfo(ctx context.Context, id int64) (res *model.Cron, err error) {
	rows, err := d.db.Query("SELECT * FROM crons where id=?", id)
	if rows == nil {
		return
	}

	res = &model.Cron{}
	for rows.Next() {
		var command string
		err = rows.Scan(&res.Id, &res.Name, &res.Status, &res.Expr, &command, &res.CreatedAt, &res.UpdatedAt)
		if err != nil {
			return
		}
		originCommand := &model.Command{}
		_ = json.Unmarshal([]byte(command), originCommand)
		res.Command = originCommand
	}
	return
}

func (d *Dao) CronList(ctx context.Context, page, pageSize int64) (res model.CronsList, err error) {

	rows, err := d.db.Query("SELECT * FROM crons order by id desc LIMIT ? OFFSET ?", pageSize, pageSize*(page-1))
	if rows == nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		info := &model.Cron{}
		var command string

		err = rows.Scan(&info.Id, &info.Name, &info.Status, &info.Expr, &command, &info.CreatedAt, &info.UpdatedAt)
		if err != nil {
			return
		}
		originCommand := &model.Command{}
		_ = json.Unmarshal([]byte(command), originCommand)
		info.Command = originCommand

		res = append(res, info)
	}
	return
}
