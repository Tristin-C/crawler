package model

import "time"

type CronsList []*Cron

type Command struct {
	Type   string `json:"type"`
	Method string `json:"method"`
	Target string `json:"target"`
}

type Cron struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Status    int64     `json:"status"`
	Expr      string    `json:"expr"`
	Command   *Command  `json:"command"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	CronStatus_normal = 1 // 正常
	CronStatus_stop   = 2 // 停用
)
