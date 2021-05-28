package main

import (
	"Go_ElasticSearch7/es05/AppInit"
	"Go_ElasticSearch7/es05/Models"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"strconv"
	"sync"
)

func main() {
	page := 1
	pagesize := 500
	wg := sync.WaitGroup{}

	//用协程加快存储
	for {
		book_list := Models.BookList{}
		db := AppInit.GetDB().Order("book_id desc").Limit(pagesize).Offset((page - 1) * pagesize).Find(&book_list)
		if db.Error != nil || len(book_list) == 0 {
			break
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := AppInit.GetEsClient()
			bulk := client.Bulk()
			for _, book := range book_list {
				req := elastic.NewBulkIndexRequest()
				req.Index("books").Id(strconv.Itoa(book.BookID)).Doc(book)
				bulk.Add(req)
			}
			rsp, err := bulk.Do(context.Background())
			if err != nil {
				log.Println(err)
			} else {
				fmt.Println(rsp)
			}

		}()
		//否则不会插入后面500条
		page = page + 1 //必须有

		wg.Wait()
	}

}
