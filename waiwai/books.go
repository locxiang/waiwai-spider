package waiwai

import (
	"encoding/json"
	"fmt"
	"github.com/lexkong/log"
	"github.com/tidwall/gjson"
	"net/http"
)

type BooksTask struct {
	req  *http.Request
	Data Books
}

type Books []Book

func (books *BooksTask) Record() error {
	books.printf()
	return nil
}

//爬虫入口
func RunEntry() error {

	//入口
	url := "https://m.tititoy2688.com/query/books?type=cartoon&paged=true&size=2&page=1&category="

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	//给一个key设定为响应的value.
	req.Header.Set("Content-Type", "application/json")


	books := new(BooksTask)
	books.req = req
	books.Data = make(Books, 0, 20)
	spider.AddTask(books)
	return nil
}

func (books *BooksTask) Marshal() ([]byte, error) {
	return json.Marshal(books)
}

type Book struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Author        string  `json:"author"`
	Description   string  `json:"description"`
	Keywords      string  `json:"keywords"`
	Type          OnSale  `json:"type"`
	CategoryID    int64   `json:"categoryId"`
	Category      string  `json:"category"`
	Status        OnSale  `json:"status"`
	FreeFlag      bool    `json:"freeFlag"`
	OnSale        OnSale  `json:"onSale"`
	CoverURL      string  `json:"coverUrl"`
	ExtensionURL  string  `json:"extensionUrl"`
	LastChapter   int64   `json:"lastChapter"`
	ChapterCount  int64   `json:"chapterCount"`
	WordCount     *int64  `json:"wordCount"`
	ReadCount     int64   `json:"readCount"`
	ChapterPoints int64   `json:"chapterPoints"`
	Recommend     bool    `json:"recommend"`
	Competitive   bool    `json:"competitive"`
	Tags          string  `json:"tags"`
	Score         float64 `json:"score"`
	VipFree       bool    `json:"vipFree"`
	Exclusive     bool    `json:"exclusive"`
	Fresh         bool    `json:"fresh"`
	H             bool    `json:"h"`
	FreeInTime    bool    `json:"freeInTime"`
}

type OnSale struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//获取所有要爬的漫画列表
func (books *BooksTask) Run() error {
	//获取内容
	body, err := spider.getContent(books.req)
	if err != nil {
		return err
	}
	str := gjson.Get(body, "list").String()

	fmt.Printf("%s\n", str)

	err = json.Unmarshal([]byte(str), &books.Data)

	return err
}

func (books *BooksTask) printf() {
	fmt.Printf("printf的地址%p count:%d\n", books.Data, len(books.Data))
	for _, book := range books.Data {
		fmt.Printf("ID:%d , 书名:%s \n", book.ID, book.Name)
	}
}

//下一步
func (books *BooksTask) Next() error {
	for _, book := range books.Data {
		//把书存下来
		AddBook(book)

		// 把书全部加入到队列
		menuUrl := fmt.Sprintf("https://m.tititoy2688.com/query/book/directory?bookId=%d", book.ID)

		req, err := http.NewRequest(http.MethodGet, menuUrl, nil)
		if err != nil {
			log.Error("http NewRequest", err)
			continue
		}
		//给一个key设定为响应的value.
		req.Header.Set("Content-Type", "application/json")



		if err := new(BookMenuTask).New(req); err != nil {
			log.Error("book_menu task new error:", err)
		}
	}
	return nil
}
