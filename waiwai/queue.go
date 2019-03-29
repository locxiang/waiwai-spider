package waiwai

import (
	"container/list"
	"fmt"
	"github.com/lexkong/log"
	"sync"
)

type SyncQueue struct {
	lock *sync.Cond
	list *list.List
}

// Create a new SyncQueue
func NewSyncQueue() *SyncQueue {
	ch := &SyncQueue{
		list: list.New(),
	}
	l := new(sync.Mutex)
	ch.lock = sync.NewCond(l)
	return ch
}
func (q *SyncQueue) poplist() (v interface{}) {
	e := q.list.Front()
	v = e.Value
	q.list.Remove(e)

	return v
}

//创建消费者 ，可以多个同时消费
func (q *SyncQueue) NewConsumer() chan interface{} {
	c := make(chan interface{})

	go func() {
		defer func() {
			if e := recover(); e != nil {
				err := fmt.Errorf("%s", e)
				log.Warnf("通道关闭", err)
			}
		}()

		for {
			c <- q.Pop()
		}

	}()

	return c
}

// 堵塞消费
func (q *SyncQueue) Pop() (v interface{}) {
	q.lock.L.Lock()
	defer q.lock.L.Unlock()

	if q.list.Len() > 0 {
		v = q.poplist()
	} else {
		q.lock.Wait()

		v = q.poplist()
	}

	return
}

// 写入数据
func (q *SyncQueue) Push(v interface{}) {
	q.lock.L.Lock()
	defer q.lock.L.Unlock()
	q.list.PushBack(v)
	q.lock.Signal()

}
