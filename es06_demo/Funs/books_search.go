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

	//bool 查询
	//GET /books/_search
	//{
	//	"query": {
	//	"bool": {
	//		"must": [
	//		{"term": {"BookPress": "人民邮电出版社"}},
	//		{"match": {"BookName": "书"}}
	//		]
	//}
	//}
	//}
	//点了天津大学出版社，点搜索会出现天津大学出版社的所有
	//组合查询
	qList := make([]elastic.Query, 0)
	if searchModel.BookName != "" { //判断书名
		//加入图书名搜索条件
		machQuery := elastic.NewMatchQuery("BookName", searchModel.BookName)
		qList = append(qList, machQuery)
	}
	//加入出版社搜索条件
	if searchModel.BookPress != "" {
		pressQuery := elastic.NewTermQuery("BookPress", searchModel.BookPress)
		qList = append(qList, pressQuery)
	}
	/*
		GET /books/_search
		{
		  "query": {
		    "bool": {
		      "must": [
		        {"term": {"BookPress": "人民邮电出版社"}},
		        {"match": {"BookName": "书"}},
		        {"range": {
		          "BookPrice1": {
		            "gte": 100,
		            "lte": 200
		          }
		        }}
		      ]
		    }
		  }
		}
	*/
	//设置价格搜索范围
	if searchModel.BookPrice1Start > 0 || searchModel.BookPrice1End > 0 {
		priceRangeQuery := elastic.NewRangeQuery("BookPrice1")
		if searchModel.BookPrice1Start > 0 {
			priceRangeQuery.Gte(searchModel.BookPrice1Start)
		}
		if searchModel.BookPrice1End > 0 {
			priceRangeQuery.Lte(searchModel.BookPrice1End)
		}
		qList = append(qList, priceRangeQuery)
	}

	//处理排序
	sortList := make([]elastic.Sorter, 0)
	{
		if searchModel.OrderSet.Score {
			sortList = append(sortList, elastic.NewScoreSort().Desc())
		}
		if searchModel.OrderSet.PriceOrder == Models.OrderByPriceAsc { //从低到高
			sortList = append(sortList, elastic.NewFieldSort("BookPrice1").Asc())
		}
		if searchModel.OrderSet.PriceOrder == Models.OrderByPriceDesc { //从高到低
			sortList = append(sortList, elastic.NewFieldSort("BookPrice1").Desc())
		}
	}

	boolMustQuery := elastic.NewBoolQuery().Must(qList...)

	rsp, err := AppInit.GetEsClient().Search().Query(boolMustQuery).SortBy(sortList...).
		From((searchModel.Current - 1) * searchModel.Size).Size(searchModel.Size). //current初始为1 所以-1
		Index("books").Do(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"result": MapToBooks(rsp), "metas": gin.H{
				"total": rsp.TotalHits(),
			},
		})
	}
}
