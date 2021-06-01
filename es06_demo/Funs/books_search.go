package Funs

import (
	"Go_ElasticSearch7/es06_demo/AppInit"
	"Go_ElasticSearch7/es06_demo/Models"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

func SeachBook(ctx *gin.Context) {
	searchModel := Models.NewSearchModel()
	err := ctx.BindJSON(searchModel)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err,
		})
		return
	}
	//书名
	machQuery := elastic.NewMatchQuery("BookName", searchModel.BookName)
	rsp, err := AppInit.GetEsClient().Search().Query(machQuery).
		Index("books").Do(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"result": MapToBooks(rsp),
		})
	}
}
