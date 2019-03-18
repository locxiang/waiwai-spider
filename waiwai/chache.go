package waiwai

import (
	"fmt"
	"sync"
)

var BooksMap = new(sync.Map)

func AddBook(book Book) {
	BooksMap.Store(book.ID, book)
}

func GetBook(bookId int64) (book Book, err error) {
	value, ok := BooksMap.Load(bookId)
	if ok {
		book, _ = value.(Book)
	} else {
		err = fmt.Errorf("bookId %d 不存在", bookId)
	}
	return
}
