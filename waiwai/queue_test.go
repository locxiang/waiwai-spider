package waiwai_test

import (
	"fmt"
	"github.com/locxiang/waiwai-spider/waiwai"
	"testing"
	"time"
)

func TestNewSyncQueue(t *testing.T) {
	q := waiwai.NewSyncQueue()

	//读
	go func() {
		consumer := q.NewConsumer()
		i := 0
		for v := range consumer {
			i++
			if i > 5 {
				close(consumer)
			}
			fmt.Printf("消费者1:  %d \n", v)
		}
	}()


	go func() {
		consumer := q.NewConsumer()
		i := 0
		for v := range consumer {
			i++
			if i > 5 {
				close(consumer)
			}
			fmt.Printf("消费者2:  %d \n", v)
		}
	}()

	time.Sleep(1 * time.Second)
	//写
	for i:=1;i<=10;i++{
		q.Push(i)
	}
	time.Sleep(1 *time.Second)
}
