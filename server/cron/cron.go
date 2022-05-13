package cron

import (
	"context"
	"crawler/conf"
	"crawler/cron"
	"crawler/service"
	"crawler/utils"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

var (
	config *conf.Config
	svc    *service.Service
)

func Run(c *conf.Config, s *service.Service) {
	config = c
	svc = s

	c2 := cron.New()
	err := cronAdd(context.Background(), c2)
	if err != nil {
		logrus.Errorf("cron start fail, err:%v", err)
		return
	}
	c2.Start()
}

func cronAdd(ctx context.Context, c *cron.Cron) error {

	crons := config.Crons
	if len(crons) <= 0 {
		return nil
	}

	errMsg := make([]string, 0)

	for _, item := range crons {
		err := c.AddFunc(item.Expr, svc.AddFunc(ctx, item))
		if err != nil {
			errM := fmt.Sprintf("Cron Name:%s, add func fail, err:%s", item.Name, err)
			logrus.Error(errM)
			errMsg = append(errMsg, errM)
		}
	}

	if len(errMsg) > 0 {
		return errors.New(utils.ToJson(errMsg))
	}
	return nil
}
