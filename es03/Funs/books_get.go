package Funs

import (
	"Go_ElasticSearch7/es03/AppInit"
	"Go_ElasticSearch7/es03/Models"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"reflect"
)

//返回rsp，，，结果集							使用指针，防止数据进行赋值
func MapToBooks(rsp *elastic.SearchResult) []*Models.Books {
	ret := []*Models.Books{}
	var t *Models.Books
	for _, item := range rsp.Each(reflect.TypeOf(t)) {
		ret = append(ret, item.(*Models.Books))
	}
	return ret
}

func LoadBook(ctx *gin.Context) {
	//显示10条 对应 GET /books/_search
	rsp, err := AppInit.GetEsClient().Search().Index("books").Do(ctx)
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
