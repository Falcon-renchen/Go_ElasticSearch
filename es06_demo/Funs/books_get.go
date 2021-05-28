package Funs

import (
	"Go_ElasticSearch7/es06_demo/AppInit"
	"Go_ElasticSearch7/es06_demo/Models"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"reflect"
	"strings"
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

func LoadBookByPress(ctx *gin.Context) {
	press, _ := ctx.Params.Get("press") // press 是  press/:press 后面的这个  必须对应
	termQuery := elastic.NewTermQuery("BookPress", press)
	rsp, err := AppInit.GetEsClient().Search().Query(termQuery).Index("books").Do(ctx)
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

func LoadBooksByPress(ctx *gin.Context) {
	press, _ := ctx.Params.Get("press") // press 是  press/:press 后面的这个  必须对应
	list := strings.Split(press, ",")
	pressList := []interface{}{}
	//循环遍历list，并添加到pressList中
	for _, p := range list {
		pressList = append(pressList, p)
	}
	termQuery := elastic.NewTermsQuery("BookPress", pressList...)
	rsp, err := AppInit.GetEsClient().Search().Query(termQuery).Index("books").Do(ctx)
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
