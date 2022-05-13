package dao

import (
	"context"
	"crawler/model"
	"crawler/utils"
	"testing"
	"time"
)

func TestStorage_BatchCron(t *testing.T) {
	cronList := make(model.CronsList, 0)
	cronList = append(cronList, &model.Cron{
		Name: "test_001",
		Expr: "* */2 * * * *",
		Command: &model.Command{
			Type:   "http",
			Method: "GET",
			Target: "www.baidu.com",
		},
		Status:    model.CronStatus_normal,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	cronList = append(cronList, &model.Cron{
		Name: "test_002",
		Expr: "* */5 * * * *",
		Command: &model.Command{
			Type:   "http",
			Method: "GET",
			Target: "www.baidu.com",
		},
		Status:    model.CronStatus_normal,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	res, err := dao.BatchCron(context.TODO(), cronList)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("test BatchCron success, ids:%v", res)
}

func TestCronsList(t *testing.T) {
	list, err := dao.CronList(context.TODO(), 1, 20)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("test CronsList success, result:%s", utils.ToJson(list))
}
