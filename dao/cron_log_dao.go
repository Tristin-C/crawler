package dao

import (
	"context"
	"crawler/model"
)

func (d *Dao) AddCronsLog(ctx context.Context, log *model.CronsLog) (err error) {

	stmt, err := d.db.Prepare("insert into crons_log(id, cron_name, exec_time, http_code, http_context_size, http_context, created_at) values(?,?,?,?,?,?,?)")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(log.Id, log.CronName, log.ExecTime, log.HttpCode, log.HttpContextSize, log.HttpContext, log.CreatedAt)
	if err != nil {
		return
	}

	return

}

func (d *Dao) CronsLogList(ctx context.Context, page, pageSize int64) (res model.CronsLogList, err error) {

	rows, err := d.db.Query("SELECT * FROM crons_log order by id desc LIMIT ? OFFSET ?", pageSize, pageSize*(page-1))
	if rows == nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		info := &model.CronsLog{}
		err = rows.Scan(&info.Id, &info.CronName, &info.ExecTime, &info.HttpCode, &info.HttpContextSize, &info.HttpContext, &info.CreatedAt)
		if err != nil {
			return
		}
		res = append(res, info)
	}
	return
}
