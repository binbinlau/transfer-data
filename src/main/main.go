package main

import (
	"flag"
	"fmt"
	"github.com/binsix/transfer-data/src/main/utils"
)

func main() {
	var appPath string
	flag.StringVar(&appPath, "app-path", utils.GetRootPath(), "111")
	flag.Parse()
	fmt.Printf("App path: %s \n", appPath)
	fmt.Printf("Conf is: %v \n", utils.Conf.Mysql.User)
}
