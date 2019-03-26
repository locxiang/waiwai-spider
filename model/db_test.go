package model

import (
	"github.com/locxiang/waiwai-spider/config"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	if err := config.Init("../config.yaml"); err != nil {
		panic(err)
	}

}

func TestConnDB(t *testing.T) {
	mysqlCfg := config.Values.Mysql
	cfg := Config{
		UserName: mysqlCfg.User,
		Password: mysqlCfg.Pass,
		Addr:     mysqlCfg.Addr,
		DbName:   mysqlCfg.DB,
	}

	type args struct {
		c Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "正常连接",
			args: args{
				c: cfg,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ConnDB(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("ConnDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
