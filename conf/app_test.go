package conf

import (
	"crawler/utils"
	"flag"
	"testing"
)

func TestInit(t *testing.T) {
	path := flag.String("conf", "/home/tenero/go-project/crawler/cmd/conf.yaml", "conf path")
	flag.Parse()
	err := Init(*path)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("config info:%s", utils.ToJson(Conf))
}
