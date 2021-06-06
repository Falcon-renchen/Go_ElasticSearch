package Middleware

import (
	"Go_ElasticSearch7/es06_demo/AppInit"
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"log"
	"strings"
)

//可以用hook 同时打印日志到es里面，不一定要输入到文件里面
type EsHook struct {
}

func NewEsHook() *EsHook {
	return &EsHook{}
}

func (this *EsHook) Fire(entry *logrus.Entry) error {
	data := entry.Data
	data["time"] = entry.Time
	data["level"] = entry.Level
	data["msg"] = entry.Message
	//开头如果有则>=0
	if strings.Index(data["url"].(string), "/favicon.ico") >= 0 {
		return nil
	}
	client := AppInit.GetEsClient()
	bulk := client.Bulk()
	{
		req := elastic.NewBulkIndexRequest()
		req.Index("bookslogs").Doc(data) //直接插入
		bulk.Add(req)
	}
	_, err := bulk.Do(context.Background())
	if err != nil {
		log.Println(err)
	}
	return nil
}

//Hook可以勾住所有类别或者想要的类别
func (this *EsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
