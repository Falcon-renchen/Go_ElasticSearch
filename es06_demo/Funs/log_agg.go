package Funs

import (
	"Go_ElasticSearch7/es06_demo/AppInit"
	"Go_ElasticSearch7/es06_demo/Middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type AggFunc func(field string) elastic.Aggregation

var app_map = map[string]AggFunc{
	"max": func(field string) elastic.Aggregation {
		return elastic.NewMaxAggregation().Field(field)
	},
}

const logIndexName = "bookslogs"

func LogAgg(ctx *gin.Context) {
	getType := ctx.Param("type")
	getFild := ctx.Param("field")
	/*
		POST /bookslogs/_search?size=0
		{
			"aggs": {
			"max_duration": {
				"max": {
					"field": "duration"
				}
			}
		}
		}
	*/
	if f, ok := app_map[getType]; ok {
		agg_name := fmt.Sprintf("%s_%s", getType, getFild)
		//自动生成agg查询
		rsp, err := AppInit.GetEsClient().Search().Aggregation(agg_name, f(getFild)).
			Index(logIndexName).Do(ctx)
		Middleware.CheckError(err, "agg request error")
		ctx.JSON(200, gin.H{
			"result": rsp.Aggregations,
		})
	} else {
		Middleware.CheckError(fmt.Errorf("agg type error"), "")
	}
}
