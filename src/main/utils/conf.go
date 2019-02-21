package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

type conf struct {
	Mysql struct {
		User       string `yaml:"user"`
		Password   string `yaml:"password"`
		Collection string `yaml:"collection"`
	} `yaml:"mysql"`
}

var Conf = GetConf()

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

// func (c *conf) Get(prop string) string {
// 	return GetConf().
// }

func GetConf() *conf {
	var resourceConf conf
	return resourceConf.getConf()
}
