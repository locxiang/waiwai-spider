package model

import (
	"github.com/jinzhu/gorm"
	"github.com/locxiang/waiwai-spider/util"
	"github.com/pborman/uuid"
	"time"
)

type IModel interface {
	Migrate() error
}

type BaseModel struct {
	Guid      string `gorm:"primary_key;unique_index" json:"guid"`
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"`
}

//
////创建的时候先处理guid
func (bm *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Guid", uuid.New())
	scope.SetColumn("CreatedAt", util.TimeMillisecond(time.Now()))
	return nil
}

func (bm *BaseModel) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", util.TimeMillisecond(time.Now()))
	return nil
}

func (bm *BaseModel) BeforeSave(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", util.TimeMillisecond(time.Now()))
	return nil
}

//用字符串存储json结构
type StrJson string

func (s StrJson) MarshalJSON() (b []byte, err error) {
	if s == "" {
		return []byte("null"), nil
	}

	return []byte(s), nil
}

type StrArray string

func (s StrArray) MarshalJSON() (b []byte, err error) {
	//
	if s == "" {
		return []byte("[]"), nil
	}

	return []byte(s), nil
}
