package model

import "time"

type CronsLogList []*CronsLog

type CronsLog struct {
	Id              int64     `json:"id"`
	CronName        string    `json:"cron_name"`
	ExecTime        time.Time `json:"exec_time"` // 任务执行时间
	HttpCode        int64     `json:"http_code"`
	HttpContextSize int64     `json:"http_context_size"`
	HttpContext     string    `json:"http_context"`
	CreatedAt       time.Time `json:"created_at"`
}
