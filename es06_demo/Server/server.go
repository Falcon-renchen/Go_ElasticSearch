package main

import (
	"Go_ElasticSearch7/es06_demo/Funs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	g := router.Group("/books")
	{
		//图书列表
		g.Handle("GET", "", Funs.LoadBook)

		//http://localhost:8080/books/press/湘潭大学出版社
		g.Handle("GET", "/press/:press", Funs.LoadBookByPress)

		//http://localhost:8080/books/presses/湘潭大学出版社,人民邮电出版社
		g.Handle("GET", "/presses/:press", Funs.LoadBooksByPress)
	}

	router.StaticFS("/ui", http.Dir("./htmls"))

	router.Run(":8080")
}
