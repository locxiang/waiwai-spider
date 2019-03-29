package waiwai

import (
	"encoding/json"
	"fmt"
	"github.com/lexkong/log"
	"github.com/locxiang/waiwai-spider/model"
	"github.com/pkg/errors"
	"net/http"
)

func (b *ChapterList) Marshal() ([]byte, error) {
	return json.Marshal(b)
}

type ChapterTask struct {
	req  *http.Request
	Data ChapterList
}

func (b *ChapterTask) New(req *http.Request) error {
	b.req = req
	b.Data = make(ChapterList, 0)
	spider.AddTask(b)
	return nil
}

func (b *ChapterTask) Run() error {
	//获取内容
	body, err := spider.getContent(b.req)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(body), &b.Data)

	return err
}

func (b *ChapterTask) Next() error {
	for _, chapter := range b.Data {
		is, err := chapter.CheckUpdate()
		if err != nil {
			log.Error("chapter CheckUpdate", err)
			continue
		}

		if is == false {
			continue
		}

		log.Infof("发现【%d】新内容【%s】", chapter.BookID, chapter.Title)

		chapterUrl := fmt.Sprintf("https://m.tititoy2688.com/query/book/chapter?bookId=%d&chapterId=%d", chapter.BookID, chapter.ID)
		req, err := http.NewRequest(http.MethodGet, chapterUrl, nil)
		if err != nil {
			log.Error("http NewRequest", err)
			continue
		}
		//给一个key设定为响应的value.
		req.Header.Set("Content-Type", "application/json")

		if err := NewDetailTask(req, chapter); err != nil {
			log.Error("book_menu task new error:", err)
		}
	}

	return nil
}

func (b *ChapterTask) Record() error {
	b.printf()
	return nil
}

//书里的所有章节
type ChapterList []Chapter

//章节描述
type Chapter struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	BookID      int64  `json:"bookId"`
	Sequence    int64  `json:"sequence"`
	CoverURL    string `json:"coverUrl"`
	FreeFlag    bool   `json:"freeFlag"`
	Points      int64  `json:"points"`
	HasRead     bool   `json:"hasRead"`
	HasPurchase bool   `json:"hasPurchase"`
	VipFree     bool   `json:"vipFree"`
	FreeInTime  bool   `json:"freeInTime"`
}

func (b *Chapter) CheckUpdate() (is bool, err error) {
	c := &model.Chapter{
		ID:          b.ID,
		Title:       b.Title,
		BookID:      b.BookID,
		Sequence:    b.Sequence,
		CoverURL:    b.CoverURL,
		FreeFlag:    b.FreeFlag,
		Points:      b.Points,
		HasRead:     b.HasRead,
		HasPurchase: b.HasPurchase,
		VipFree:     b.VipFree,
		FreeInTime:  b.FreeInTime,
	}

	found, err := new(model.Chapter).Get(b.ID)
	if err != nil && found == false {
		// 有错误，并且不是数据不存在
		return false, errors.Wrap(err, "获取信息失败")
	}

	if found {
		err := c.Create()
		if err != nil {
			return false, errors.Wrap(err, "创建Chapter数据信息失败")
		}

		return true, nil
	}

	return false, nil

}

func (b *ChapterTask) printf() {
	log.Infof("检索%d个章节", len(b.Data))
}
