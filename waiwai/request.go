package waiwai

import (
	"context"
	"fmt"
	"github.com/lexkong/log"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
	"time"
)

//单例爬虫
var spider *Spider

//任务接口
type Tasker interface {
	//New(req *http.Request) error //创建任务
	Run() error    //执行任务
	Next() error   //获取到内容后的下一步
	Record() error //用于记录此次任务的情况
}

//定义爬虫
type Spider struct {
	tasks      *SyncQueue         //执行任务队列
	Sleep      func()             //暂停时间
	client     *http.Client       //http 引擎
	cancel     context.CancelFunc //任务完全结束
	concurrent int                //并发数
}

//http client 引擎
func (r *Spider) request() *http.Client {
	return r.client
}

//创建一个爬虫
func New(ctx context.Context, cancel context.CancelFunc) *Spider {
	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar:     jar,
		Timeout: 100 * time.Second,
	}

	spider = &Spider{
		tasks: NewSyncQueue(),
		Sleep: func() {
			log.Debugf("执行暂停")
			time.Sleep(200 * time.Millisecond) //暂停时间
		},
		client:     client,
		cancel:     cancel,
		concurrent: 1,
	}

	for i := spider.concurrent; i > 0; i-- {
		go spider.executeTask(ctx)
	}
	//开启任务执行
	return spider
}

//添加任务
func (r *Spider) AddTask(task Tasker) {
	r.tasks.Push(task)
}

//执行任务
func (r *Spider) executeTask(ctx context.Context) {
	c := spider.tasks.NewConsumer()
	for {
		select {
		case <-ctx.Done():
			log.Infof("ctx.Done 结束")
			return

		case <-time.After(10 * time.Second):
			log.Debugf("暂时没有任务，结束")
			r.cancel()
			return
		case v := <-c:
			task := v.(Tasker)
			//执行任务
			for i, p := 1, 1; true; i *= 2 {
				if err := task.Run(); err != nil {
					log.Error("task run ", err)
					for j := i; j > 0; j-- {
						r.Sleep()
					}
					log.Infof("第%d次重试：...", p)
					p++
					continue
				} else {
					break
				}
			}

			//记录任务情况
			if err := task.Record(); err != nil {
				log.Error("task record ", err)
				return
			}

			//执行任务下一步
			if err := task.Next(); err != nil {
				log.Error("task next", err)
				return
			}

			//暂停
			r.Sleep()
		}
	}

}

//下载文件
func (r *Spider) downFile(req *http.Request, desFile string) error {
	log.Debugf("准备下载:%s 到 %s", req.URL, desFile)

	dir, _ := filepath.Split(desFile)
	//判断是不是有文件夹，如果没有就创建
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("client do error :%s", err)
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	f, err := os.Create(desFile)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	io.Copy(f, resp.Body)
	return nil
}

//获取内容
func (r *Spider) getContent(req *http.Request) (string, error) {
	resp, err := r.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("client do error :%s", err)
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ioutil ReadAll:%s", err)
	}

	code := gjson.GetBytes(body, "code").Int()
	msg := gjson.GetBytes(body, "code").String()
	if code != 200 {
		return "", fmt.Errorf("返回json错误:%s", msg)
	}
	//ffmt.Puts(req.Cookies, resp.Request.Cookies())

	return gjson.GetBytes(body, "content").String(), nil
}
