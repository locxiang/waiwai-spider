package model

import "github.com/pkg/errors"

type Detail struct {
	ID        int64  `json:"id"`
	BookId    int64  `json:"book_id"`
	ChapterID int64  `json:"chapter_id"`
	URL       string `json:"url"`
	Sequence  string `json:"sequence"`
	Width     int64  `json:"width"`
	Height    int64  `json:"height"`
}

func (d *Detail) Get(id int64) (found bool, err error) {
	query := DB.Model(d).Where("id = ?", id).Take(d)
	found = query.RecordNotFound()
	return found, query.Error
}

func (d *Detail) Create() error {
	query := DB.Model(d).Where("id = ?", d.ID).Take(d)
	if query.RecordNotFound() == false {
		return errors.Wrap(query.Error,"记录已经存在")
	}

	query = DB.Model(d).Create(d)
	return query.Error
}

//数据库表初始化,含写入数据
func (d *Detail) Migrate() error {
	query := DB.AutoMigrate(&d)
	return query.Error
}
