package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
	"github.com/locxiang/waiwai-spider/util"
	"time"
)

type BaseModel struct {
	Guid      string        `gorm:"primary_key;unique_index" json:"guid"`
	CreatedAt sql.NullInt64 `json:"created_at"`
	UpdatedAt sql.NullInt64 `json:"updated_at"`
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
