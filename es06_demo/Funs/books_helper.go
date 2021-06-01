package Funs

import (
	"Go_ElasticSearch7/es06_demo/AppInit"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

//取"出版社"列表作为搜索条件
func MapFiledsToSlice(rsp *elastic.SearchResult, key string) []interface{} {
	ret := make([]interface{}, 0)
	for _, hit := range rsp.Hits.Hits {
		ret = append(ret, hit.Fields[key].([]interface{})[0]) //取出Field下的第一个
	}
	return ret
}

//取前10个出版社,折叠
func PressList(ctx *gin.Context) {
	//GET /books/_search
	//{
	//	"_source": "",    //_source为空, FetchSource=false
	//	"collapse": {
	//	"field": "BookPress"
	//}
	//}
	cb := elastic.NewCollapseBuilder("BookPress")
	//默认显示10条 对应 GET /books/_search     Size(20)显示20条
	rsp, err := AppInit.GetEsClient().Search().
		Collapse(cb).FetchSource(false).Index("books").Size(20).
		Do(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err,
		})
	} else {
		//ctx.JSON(200, gin.H{
		//	"result": rsp,
		//})
		ctx.JSON(200, gin.H{
			//取出版社
			"result": MapFiledsToSlice(rsp, "BookPress"),
		})
	}
}
