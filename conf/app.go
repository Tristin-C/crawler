package conf

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Http   *Http       `yaml:"http"`
	Sqlite *Sqlite     `yaml:"sqlite"`
	Crons  []*CronItem `yaml:"crons"`
}

type Http struct {
	Addr string `yaml:"addr"`
}

type Sqlite struct {
	Dsn string `yaml:"dsn"`
}

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

var (
	Conf *Config
)

func Init(path string) (err error) {

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Errorf("conf.Init ioutil.ReadFile fail, err:%s", err)
		return err
	}

	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		logrus.Errorf("conf.Init yaml.Unmarshal fail, err:%s", err)
		return err
	}

	return nil
}
