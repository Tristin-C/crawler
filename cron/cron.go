package cron

import (
	crontab "github.com/robfig/cron"
)

type Cron struct {
	cron *crontab.Cron
}

func New() *Cron {
	c := &Cron{
		cron: crontab.New(),
	}
	return c
}

func (c *Cron) AddFunc(spec string, cmd func()) error {
	return c.cron.AddFunc(spec, cmd)
}

func (c *Cron) AddJob(spec string, cmd crontab.Job) error {
	return c.cron.AddJob(spec, cmd)
}

func (c *Cron) Start() {
	c.cron.Start()
	select {}
}
