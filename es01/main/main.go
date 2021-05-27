package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

func main() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://172.16.17.156:9200/"),
		elastic.SetSniff(false), //本机地址用true,虚拟机地址用false
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background() //假设不会超时
	//mapping, err := client.GetMapping().Index("news").Do(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(mapping)

	json := `{"news_title":"test1","news_type":"php","news_status":1}`

	//获取索引，类似数据库表
	data, err := client.Index().Index("news").Id("101").BodyString(json).Do(ctx)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}
