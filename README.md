# 歪歪漫画爬虫

## 爬取流程

1. 获取所有要爬的漫画数据列表
2. 获取漫画目录详情
3. 获取漫画详情
4. 下载整理详情的内容
	
## 注意事项
1. ip不能过于频繁
2. 没有登录无法获取漫画之后的详情

## 技术思路 （结构体）

### Spider 

创建一个spider 结构体
存放这个网站的相关采集结构

### Tasker

创建一个task 的结构体用于存储每次http的相关数据，以便于控制并发和cookie管理
每个任务都是3个步骤

1. 执行
2. 记录/报告 情况
3. 下一步

### Queue
 并发安全的内存队列，支持多消费

## 采集分析说明


获取所有的漫画清单（已经下载在本地）
https://m.tititoy2688.com/query/books?type=cartoon&paged=true&size=2000&page=1&category=


书籍目录
https://m.tititoy2688.com/query/book/directory?bookId=1296


书籍详情
https://m.tititoy2688.com/query/book/chapter?bookId=1296&chapterId=226092



没有token获取不到书单，但是可以看图片内容


换IP访问可以增加金币
https://m.tititoy2688.com/pages/?inviter=786518


JSESSIONID  token需要注意，一直在变化，需要做一个策略作为刷新
