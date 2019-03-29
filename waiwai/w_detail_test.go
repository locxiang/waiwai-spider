package waiwai

import (
	"github.com/locxiang/waiwai-spider/config"
)

func init() {
	// 初始化配置文件
	if err := config.Init("../config.yaml"); err != nil {
		panic(err)
	}
}