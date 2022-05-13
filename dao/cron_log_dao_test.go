package dao

import (
	"context"
	"crawler/model"
	"crawler/utils"
	"testing"
	"time"
)

func TestAddCronsLog(t *testing.T) {
	err := dao.AddCronsLog(context.TODO(), &model.CronsLog{
		CronName:        "test_1",
		ExecTime:        time.Now(),
		HttpCode:        200,
		HttpContextSize: 1024,
		HttpContext:     "success",
		CreatedAt:       time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("test addCronLog success")
}

func TestCronsLogList(t *testing.T) {
	list, err := dao.CronsLogList(context.TODO(), 1, 20)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("test cronsLogList success, result:%s", utils.ToJson(list))
}
