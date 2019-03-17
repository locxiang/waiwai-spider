package main

import (
	"github.com/locxiang/waiwai-spider/waiwai"
	"github.com/spf13/pflag"
	"runtime"
	"github.com/locxiang/waiwai-spider/config"
)

var (
	cfg = pflag.StringP("config", "c", "", "config file path.")
)

func main() {
	//使用全部cpu
	runtime.GOMAXPROCS(runtime.NumCPU())
	pflag.Parse()

	// 初始化配置文件
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	waiwai.New()

	waiwai.RunEntry()

	select {}
}
