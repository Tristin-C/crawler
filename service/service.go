package service

import (
	"context"
	"crawler/conf"
	"crawler/dao"
	"crawler/model"
	"crawler/utils"
)

type Service struct {
	c     *conf.Config
	st    *dao.Dao
	idGen *utils.IdGenService
}

func NewService(c *conf.Config) (s *Service) {
	idGen := utils.NewIdGenService()
	s = &Service{
		c:     c,
		st:    dao.NewDao(c, idGen),
		idGen: idGen,
	}
	return s
}

func (s *Service) Close() {
	s.st.Close()
}

// crons
func (s *Service) BatchCron(ctx context.Context, cronList model.CronsList) (ids []int64, err error) {
	ids, err = s.st.BatchCron(ctx, cronList)
	if err != nil {
		return
	}

	for _, item := range cronList {
		s.AddFunc(ctx, &conf.CronItem{
			Expr:    item.Expr,
			Name:    item.Name,
			Command: (*conf.Command)(item.Command),
		})
	}
	return
}

func (s *Service) CronInfo(ctx context.Context, id int64) (res *model.Cron, err error) {
	return s.st.CronInfo(ctx, id)
}

func (s *Service) CronsList(ctx context.Context, page, pageSize int64) (res model.CronsList, err error) {
	return s.st.CronList(ctx, page, pageSize)
}

// crons_log
func (s *Service) AddCronsLog(ctx context.Context, log *model.CronsLog) (err error) {
	return s.st.AddCronsLog(ctx, log)
}

func (s *Service) CronsLogList(ctx context.Context, page, pageSize int64) (res model.CronsLogList, err error) {
	return s.st.CronsLogList(ctx, page, pageSize)
}
