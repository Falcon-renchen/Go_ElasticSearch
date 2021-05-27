package main

import (
	"Go_ElasticSearch7/es03/Funs"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	g := router.Group("/books")
	{
		g.Handle("GET", "", Funs.LoadBook)
	}
	router.Run(":8080")
}
