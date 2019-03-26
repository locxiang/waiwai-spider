package main

import (
	"context"
	"github.com/lexkong/log"
	"github.com/locxiang/waiwai-spider/config"
	"github.com/locxiang/waiwai-spider/model"
	"github.com/locxiang/waiwai-spider/waiwai"
	"github.com/spf13/pflag"
	"runtime"
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
	mysqlCfg := config.Values.Mysql
	cfg := model.Config{
		UserName: mysqlCfg.User,
		Password: mysqlCfg.Pass,
		Addr:     mysqlCfg.Addr,
		DbName:   mysqlCfg.DB,
	}
	if err := model.ConnDB(cfg); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	waiwai.New(ctx, cancel)

	waiwai.RunEntry()

	<-ctx.Done()

	log.Info("完成")

}
