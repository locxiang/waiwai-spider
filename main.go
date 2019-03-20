package main

import (
	"context"
	"github.com/lexkong/log"
	"github.com/locxiang/waiwai-spider/config"
	"github.com/locxiang/waiwai-spider/waiwai"
	"github.com/spf13/pflag"
	"os"
	"runtime"
	"runtime/trace"
)

var (
	cfg = pflag.StringP("config", "c", "", "config file path.")
)

func main() {


	f, _ := os.Create("trace.out")
	trace.Start(f)

	defer func() {
		f.Close()
		trace.Stop()
	}()


	//使用全部cpu
	runtime.GOMAXPROCS(runtime.NumCPU())
	pflag.Parse()

	// 初始化配置文件
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	waiwai.New(ctx, cancel)


	waiwai.RunEntry()

	<-ctx.Done()

	log.Info("完成")

}
