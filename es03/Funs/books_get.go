package Funs

import (
	"Go_ElasticSearch7/es03/AppInit"
	"github.com/gin-gonic/gin"
)

func LoadBook(ctx *gin.Context) {
	//显示10条 对应 GET /books/_search
	rsp, err := AppInit.GetEsClient().Search().Index("books").Do(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"result": rsp.Hits.Hits,
		})
	}
}
