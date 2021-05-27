package main

import (
	"Go_ElasticSearch7/es01/AppInit"
	"Go_ElasticSearch7/es01/Models"
	"context"
	"fmt"
	"strconv"
)

func main() {
	page := 1
	pagesize := 500

	for {
		book_list := Models.BookList{}
		db := AppInit.GetDB().Order("book_id desc").Limit(pagesize).Offset((page - 1) * pagesize).Find(&book_list)
		if db.Error != nil || len(book_list) == 0 {
			break
		}
		//将50条信息添加到kb里面  通过GET /books/_count 查看
		for _, book := range book_list {
			ctx := context.Background()
			rsp, err := AppInit.GetEsClient().Index().Index("books").
				Id(strconv.Itoa(book.BookID)).BodyJson(book).Do(ctx)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(rsp)
			}
		}

		break
	}

}
