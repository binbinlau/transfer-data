package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index]
}

func GetConfPath(filedir string, filename string) string {
	var err error
	var appPath, appConfigPath string
	if filedir == "" {
		filedir = "src/main/resource" //默认配置文件路径
	}
	if appPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	}
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appConfigPath = filepath.Join(workPath, filedir, filename)
	fmt.Printf("workPath: %s \n", appConfigPath)
	if !FileExists(appConfigPath) {
		appConfigPath = filepath.Join(appPath, filedir, filename)
	}
	return appConfigPath
}
