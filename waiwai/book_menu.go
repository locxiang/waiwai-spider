package waiwai

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/lexkong/log"
)

func (b *BookMenu) Marshal() ([]byte, error) {
	return json.Marshal(b)
}

type BookMenuTask struct {
	req  *http.Request
	Data BookMenu
}

func (b *BookMenuTask) New(req *http.Request) error {
	b.req = req
	b.Data = make(BookMenu, 0)
	spider.AddTask(b)
	return nil
}

func (b *BookMenuTask) Run() error {
	//获取内容
	body, err := spider.getContent(b.req)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(body), &b.Data)

	return err
}

func (b *BookMenuTask) Next() error {
	for _, bookChapter := range b.Data {

		bookChapterUrl := fmt.Sprintf("https://m.tititoy2688.com/query/book/chapter?bookId=%d&chapterId=%d", bookChapter.BookID, bookChapter.ID)
		req, err := http.NewRequest(http.MethodGet, bookChapterUrl, nil)
		if err != nil {
			log.Error("http NewRequest", err)
			continue
		}
		//给一个key设定为响应的value.
		req.Header.Set("Content-Type", "application/json")

		if err := NewBookDetailsTask(req, bookChapter); err != nil {
			log.Error("book_menu task new error:", err)
		}
	}

	return nil
}

func (b *BookMenuTask) Record() error {
	b.printf()
	return nil
}

//书里的所有章节
type BookMenu []BookChapter

//章节描述
type BookChapter struct {
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

func (b *BookMenuTask) printf() {
	for _, m := range b.Data {
		fmt.Printf("标题:%s \n", m.Title)
	}
}
