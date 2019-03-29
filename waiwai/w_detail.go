package waiwai

import (
	"encoding/json"
	"fmt"
	"github.com/lexkong/log"
	"github.com/locxiang/waiwai-spider/model"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"net/http"
	"strconv"
)

//获取章节详情
type DetailTask struct {
	req     *http.Request
	Chapter Chapter
	Data    DetailList
}

func NewDetailTask(req *http.Request, Chapter Chapter) error {
	b := new(DetailTask)
	b.req = req
	b.Chapter = Chapter
	b.Data = make(DetailList, 0)
	spider.AddTask(b)
	return nil
}

func (b *DetailTask) Run() error {
	log.Debugf("准备采集：%s", b.req.URL)

	//获取内容
	body, err := spider.getContent(b.req)
	if err != nil {
		return err
	}

	body = gjson.Get(body, "imageList").String()
	if body == "" {
		return errors.New("imageList is null")
	}
	err = json.Unmarshal([]byte(body), &b.Data)

	//写入附加信息
	for i := range b.Data {
		b.Data[i].BookId = b.Chapter.BookID
		b.Data[i].ChapterID = b.Chapter.ID
	}

	return err
}

func (b *DetailTask) Next() error {
	//下载图片
	for _, detail := range b.Data {

		is, err := detail.CheckUpdate()
		if err != nil {
			log.Error("detail CheckUpdate", err)
			continue
		}

		if is == false {
			continue
		}

		req, err := http.NewRequest(http.MethodGet, detail.URL, nil)
		if err != nil {
			log.Error("http NewRequest", err)
			continue
		}

		book, err := GetBook(b.Chapter.BookID)
		if err != nil {
			log.Error("BookDetailsTask Next GetBook", err)
			continue
		}

		filename := fmt.Sprintf("%s/%s/%s.jpg", book.Name, b.Chapter.Title, detail.ID)
		spider.downFile(req, "/tmp/books/"+filename)
	}

	return nil
}

func (b *DetailTask) Record() error {
	log.Debugf("发现图片数量:%d", len(b.Data))
	return nil
}

type DetailList []Detail

func (r *DetailList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Detail struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	Sequence string `json:"sequence"`
	Width    int64  `json:"width"`
	Height   int64  `json:"height"`
	//附加信息
	BookId    int64 `json:"book_id"`
	ChapterID int64 `json:"chapter_id"`
}

func (d *Detail) CheckUpdate() (is bool, err error) {
	detailId, err := strconv.ParseInt(d.ID, 10, 64)
	if err != nil {
		return false, errors.Wrap(err, "detailId转换失败")
	}
	m := &model.Detail{
		ID:        detailId,
		URL:       d.URL,
		Sequence:  d.Sequence,
		Width:     d.Width,
		Height:    d.Height,
		BookId:    d.BookId,
		ChapterID: d.ChapterID,
	}

	found, err := new(model.Chapter).Get(detailId)
	if err != nil && found == false {
		// 有错误，并且不是数据不存在
		return false, errors.Wrap(err, "获取Detail信息失败")
	}

	if found {
		err := m.Create()
		if err != nil {
			return false, errors.Wrap(err, "创建Detail数据信息失败")
		}
		return true, nil
	}

	return false, nil
}
