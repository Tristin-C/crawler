package service

import (
	"context"
	"crawler/conf"
	"crawler/model"
	"crawler/utils"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/sirupsen/logrus"
)

type CronItem struct {
	Expr    string   `yaml:"expr"`
	Name    string   `yaml:"name"`
	Command *Command `yaml:"command"`
}

type Command struct {
	Type   string `yaml:"type"`
	Method string `yaml:"method"`
	Target string `yaml:"target"`
}

func (s *Service) AddFunc(ctx context.Context, c *conf.CronItem) func() {
	switch c.Command.Type {
	case "http", "https":
		return s.httpFunc(ctx, c)
	default:
		return nil
	}
}

func (s *Service) httpFunc(ctx context.Context, c *conf.CronItem) func() {
	switch c.Command.Method {
	case "GET":
		return s.httpGetFunc(ctx, c)
	default:
		return nil
	}
}

func (s *Service) httpGetFunc(ctx context.Context, c *conf.CronItem) func() {
	return func() {
		err := s.ExecGetFunc(ctx, c)
		if err != nil {
			logrus.Errorf("cron name: %s exec result is failed, cron info: %s err: %s", c.Name, utils.ToJson(c), err)
			return
		}
		logrus.Infof("cron name: %s exec success, cron info: %s", c.Name, utils.ToJson(c))
	}
}

func (s *Service) ExecGetFunc(ctx context.Context, c *conf.CronItem) error {
	header := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HttpGet(ctx, c.Command.Target, nil, header)
	if err != nil {
		return err
	}
	if resp == nil {
		return fmt.Errorf("resp is nil, request info: %s", utils.ToJson(c))
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var (
		httpCode        int64  = int64(resp.StatusCode)
		httpContextSize int64  = resp.ContentLength
		httpContext     string = string(respBody)
	)

	// 记录日志
	err = s.st.AddCronsLog(ctx, &model.CronsLog{
		Id:              s.idGen.GetId(),
		CronName:        c.Name,
		ExecTime:        time.Now(),
		HttpCode:        httpCode,
		HttpContextSize: httpContextSize,
		HttpContext:     httpContext,
		CreatedAt:       time.Now(),
	})
	if err != nil {
		return fmt.Errorf("AddCronsLog fail, err:%v", err)
	}

	return nil
}
