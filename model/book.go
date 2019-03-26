package model

import (
	"github.com/pkg/errors"
)

type Book struct {
	ID           int64  `gorm:"primary_key;unique_index" json:"id"`        //id
	Name         string `json:"name"`                                      //书名
	Author       string `json:"author"`                                    //作者
	Description  string `json:"description"`                               //描述
	ExtensionURL string `gorm:"column:extension_url" json:"extension_url"` //封面
	Keywords     string `json:"keywords"`                                  //关键字
	Category     string `json:"category"`                                  //分类
	LastChapter  int64  `gorm:"column:last_chapter" json:"last_chapter"`   //最新章节
	ChapterCount int64  `gorm:"column:chapter_count" json:"chapter_count"` //章节数
	Tags         string `json:"tags"`                                      //标签
	Status       string `json:"status"`                                    //状态  是否完结
}

func (b *Book) Create() error {
	m := new(Book)
	query := DB.Model(m).Where("id = ?", b.ID).Take(m)
	if query.RecordNotFound() == false {
		return errors.New("记录已存在")
	}

	query = DB.Model(b).Create(b)
	return query.Error
}

/** 获取书籍信息
@param id    书籍id
@return *Book  书籍信息
@return bool  数据不存在
@return error  错误信息
 */
func (b *Book) Get(id int64) (z *Book, found bool, err error) {
	z = new(Book)
	query := DB.Model(z).Where("id = ?", id).Take(z)
	found = query.RecordNotFound()
	return z, found, query.Error
}

//数据库表初始化,含写入数据
func (b *Book) Migrate() error {
	query := DB.AutoMigrate(&b)
	return query.Error
}
