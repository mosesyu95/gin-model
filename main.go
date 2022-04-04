// @Time : 2021/4/21 下午5:29
// @Author : yuchangfu
// @File : main.go
package main

import (
	"cmdb/config"
	l "cmdb/log"
	"cmdb/model"
	"flag"
	"github.com/toolkits/pkg/runner"
	"log"
	"os"
)

var (
	help *bool
	conf *string
)

// 在main前执行，显示文件和打印版本信息
func init() {
	help = flag.Bool("h", false, "print this help.")
	conf = flag.String("f", "", "specify configuration file.")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	runner.Init()
	log.Println("runner.cwd", runner.Cwd)
	log.Println("runner.hostname", runner.Hostname)
}

func main() {
	config.Init(*conf) // 初始化配置
	model.InitDB()
	l.Init()
}
