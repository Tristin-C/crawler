package dao

import (
	"crawler/conf"
	"crawler/utils"
	"flag"
	"fmt"
	"testing"
)

var (
	dao *Dao
)

func TestMain(m *testing.M) {
	path := flag.String("conf", "/home/tenero/go-project/crawler/cmd/conf.yaml", "conf path")
	flag.Parse()
	err := conf.Init(*path)
	if err != nil {
		panic(err)
	}
	fmt.Println("conf.Init ok...")

	dao = NewDao(conf.Conf, utils.NewIdGenService())
	m.Run()
}
