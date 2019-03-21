package waiwai

import (
	"github.com/locxiang/waiwai-spider/config"
)

func init() {
	// 初始化配置文件
	if err := config.Init("../config.yaml"); err != nil {
		panic(err)
	}
}
//
//func TestBookDetailsTask_Run(t *testing.T) {
//	bookChapter := &BookChapter{
//		ID:     222279,
//		BookID: 1281,
//	}
//	bookChapterUrl := fmt.Sprintf("https://m.tititoy2688.com/query/book/chapter?bookId=%d&chapterId=%d", bookChapter.BookID, bookChapter.ID)
//	req, err := http.NewRequest(http.MethodGet, bookChapterUrl, nil)
//	if err != nil {
//		t.Errorf("http NewRequest:%s", err)
//		return
//	}
//	//给一个key设定为响应的value.
//	req.Header.Set("Content-Type", "application/json")
//
//	New()
//
//
//	if err :=NewBookDetailsTask(req,bookChapter) ; err != nil {
//		t.Error(err)
//	}
//
//	select {}
//
//}
