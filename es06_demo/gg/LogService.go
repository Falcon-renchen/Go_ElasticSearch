package gg

import (
	"Go_ElasticSearch7/es06_demo/AppInit"
	"context"
	"github.com/olivere/elastic/v7"
	"reflect"
)

func MapToLogs(rsp *elastic.SearchResult) []*LogModel {
	ret := []*LogModel{}
	var t *LogModel
	for _, item := range rsp.Each(reflect.TypeOf(t)) {
		ret = append(ret, item.(*LogModel))
	}
	return ret
}

type LogService struct {
}

func NewLogService() *LogService {
	return &LogService{}
}

//可以传参数，也可以不传参数   单参数
func (this *LogService) Search(url string) ([]*LogModel, error) {
	//不传参数，不用写这一行
	urlQuery := elastic.NewWildcardQuery("url.keyword", url)
	rsp, err := AppInit.GetEsClient().Search().Index("bookslogs").
		Query(urlQuery).
		Do(context.Background())
	return MapToLogs(rsp), err
}

//多参数
func (this *LogService) Searchs(url interface{}, method interface{}) ([]*LogModel, error) {
	qList := make([]elastic.Query, 0)
	if url != nil {
		urlQuery := elastic.NewWildcardQuery("url.keyword", url.(string))
		qList = append(qList, urlQuery)
	}

	if method != nil {
		methodQuery := elastic.NewWildcardQuery("method", method.(string))
		qList = append(qList, methodQuery)
	}
	boolMustQuery := elastic.NewBoolQuery().Must(qList...)
	rsp, err := AppInit.GetEsClient().Search().Index("bookslogs").
		//Query(urlQuery).  单参数
		Query(boolMustQuery). // 多参数
		Do(context.Background())
	return MapToLogs(rsp), err
}
