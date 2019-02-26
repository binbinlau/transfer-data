package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

type conf struct {
	Mysql struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
		Charset  string `yaml: "charset"`
	} `yaml:"mysql"`

	Logger struct {
		ShowSQL bool   `yaml: "showsql"`
		Level   string `yaml: "level"`
	} `yaml:"logger"`

	MysqlMapperDirPath string `yaml: "mysqlMapperDirPath"`
}

func (c *conf) getConf() *conf {
	rootdir := GetRootPath()
	yamlFile, err := ioutil.ReadFile(filepath.Join(rootdir, "src/main/resource/conf.yaml")) //所有的路径都是相对项目根路径
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

var Conf *conf

func init() {
	Conf = getConf()
	fmt.Println("conf is %v", Conf)
}

func getConf() *conf {
	var resourceConf conf
	return resourceConf.getConf()
}
