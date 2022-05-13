package main

import (
	"crawler/conf"
	"crawler/server/cron"
	"crawler/server/web"
	"crawler/service"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	path := flag.String("conf", "conf.yml", "conf path")
	flag.Parse()
	err := conf.Init(*path)
	if err != nil {
		panic(err)
	}
	fmt.Println("conf.Init ok...")

	svc := service.NewService(conf.Conf)
	go cron.Run(conf.Conf, svc)
	err = web.Start(conf.Conf, svc)
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	fmt.Println("server start...")
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			svc.Close()
			fmt.Println("server exit...")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
