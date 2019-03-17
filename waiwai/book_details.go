package waiwai

import (
	"encoding/json"
	"net/http"
	"github.com/lexkong/log"
	"fmt"
	"github.com/tidwall/gjson"
)

type BookDetailsTask struct {
	req     *http.Request
	Chapter *BookChapter
	Data    BookDetailsList
}

func NewBookDetailsTask(req *http.Request, Chapter *BookChapter) error {
	b := new(BookDetailsTask)
	b.req = req
	b.Chapter = Chapter
	b.Data = make(BookDetailsList, 0)
	spider.AddTask(b)
	return nil
}

func (b *BookDetailsTask) Run() error {
	log.Debugf("准备采集：%s", b.req.URL)

	//获取内容
	body, err := spider.getContent(b.req)
	if err != nil {
		return err
	}

	body = gjson.Get(body,"list").String()
	fmt.Printf("%s \n", body)
	err = json.Unmarshal([]byte(body), &b.Data)
	return err
}

func (b *BookDetailsTask) Next() error {
	log.Info("完成")

	//下载图片
	for _, detail := range b.Data {
		req, err := http.NewRequest(http.MethodGet, detail.URL, nil)
		if err != nil {
			log.Error("http NewRequest", err)
			continue
		}
		filename := fmt.Sprintf("%d_%d_%s.jpg", b.Chapter.BookID, b.Chapter.ID, detail.ID)
		spider.downFile(req, "/tmp/"+filename)
	}

	return nil
}

func (b *BookDetailsTask) Record() error {
	log.Debugf("文章数量:%d", len(b.Data))
	return nil
}

type BookDetailsList []BookDetails

func (r *BookDetailsList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type BookDetails struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	Sequence string `json:"sequence"`
	Width    int64  `json:"width"`
	Height   int64  `json:"height"`
}

//获取章节详情
func (r *Spider) GetBookChapter(bookId, chapterId int64) {

}
