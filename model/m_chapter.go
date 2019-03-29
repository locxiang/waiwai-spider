package model

import "github.com/pkg/errors"

type Chapter struct {
	ID          int64  `json:"id"`       //章节id
	Title       string `json:"title"`    //标题
	BookID      int64  `json:"bookId"`   //书id
	Sequence    int64  `json:"sequence"` //序列 1,2,3,4
	CoverURL    string `json:"coverUrl"` //封面
	FreeFlag    bool   `json:"freeFlag"`
	Points      int64  `json:"points"`
	HasRead     bool   `json:"hasRead"`
	HasPurchase bool   `json:"hasPurchase"`
	VipFree     bool   `json:"vipFree"` //vip免费
	FreeInTime  bool   `json:"freeInTime"`
}

/** 获取书籍信息
@param id    id
@return *Book  书籍信息
@return bool  数据不存在
@return error  错误信息
 */
func (c *Chapter) Get(id int64) (found bool, err error) {
	query := DB.Model(c).Where("id = ?", id).Take(c)
	found = query.RecordNotFound()
	return found, query.Error
}

func (c *Chapter) Create() error {
	query := DB.Model(c).Where("id = ?", c.ID).Take(c)
	if query.RecordNotFound() == false {
		return errors.New("记录已存在")
	}

	query = DB.Model(c).Create(c)
	return query.Error
}

func (c *Chapter) Update() error {
	query := DB.Model(c).Where("id = ?", c.ID).Update(c)
	return query.Error
}

//数据库表初始化,含写入数据
func (c *Chapter) Migrate() error {
	query := DB.AutoMigrate(&c)
	return query.Error
}
