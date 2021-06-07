package main

import (
	"Go_ElasticSearch7/es06_demo/logparser"
)

func main() {
	//将文件中日志插入到 kibana中
	p := logparser.NewHttpdParser()
	p.ParseToEs("es06_demo/logparser/logs/apache_log_2.txt")
	//list := p.Parse()
	//client := AppInit.GetEsClient()
	//bulk := client.Bulk()
	//for _, m := range list {
	//	req := elastic.NewBulkIndexRequest()
	//	req.Index("bookslogs").Doc(m) //直接插入
	//	bulk.Add(req)
	//}
	//_, err := bulk.Do(context.Background())
	//if err != nil {
	//	log.Println(err)
	//}
}
