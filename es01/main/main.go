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
	mapping, err := client.GetMapping().Index("news").Do(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mapping)
}
