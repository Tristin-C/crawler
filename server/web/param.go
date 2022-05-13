package web

import (
	"crawler/model"
	"time"
)

type CronAddRequest struct {
	List []*CronAddRequestItem `json:"list"`
}

type CronAddRequestItem struct {
	Name    string   `json:"name"`
	Expr    string   `json:"expr"`
	Command *Command `json:"command"`
}

type Command struct {
	Type   string `json:"type"`
	Method string `json:"method"`
	Target string `json:"target"`
}

func (req *CronAddRequest) ConvertToCronAddReq() (list model.CronsList) {

	if len(req.List) <= 0 {
		return
	}

	for _, item := range req.List {
		list = append(list, &model.Cron{
			Name:   item.Name,
			Status: model.CronStatus_normal,
			Expr:   item.Expr,
			Command: &model.Command{
				Type:   item.Command.Type,
				Method: item.Command.Method,
				Target: item.Command.Target,
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}
	return
}
