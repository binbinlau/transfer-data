package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type conf struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile(GetConfPath("", "conf.yaml"))
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func GetConf() *conf {
	var resourceConf conf
	return resourceConf.getConf()
}
