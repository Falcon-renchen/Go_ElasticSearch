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

//为了不顺序的输入
type LogService struct {
	quertList []elastic.Query
	size      int //显示的条数
}

func NewLogService() *LogService {
	return &LogService{make([]elastic.Query, 0), 10}
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

func (this *LogService) WhitUrlQuery(url interface{}) *LogService {
	if url != nil {
		urlQuery := elastic.NewWildcardQuery("url.keyword", url.(string))
		this.quertList = append(this.quertList, urlQuery)
	}
	return this
}
func (this *LogService) WhitMethodQuery(method interface{}) *LogService {
	if method != nil {
		methodQuery := elastic.NewWildcardQuery("url.keyword", method.(string))
		this.quertList = append(this.quertList, methodQuery)
	}
	return this
}

//不设置的话默认显示10条
func (this *LogService) WhitSize(size interface{}) *LogService {
	if size != nil {
		this.size = size.(int)
	}
	return this
}

//多参数
//通过灵活顺序的传数据
func (this *LogService) Searchs() ([]*LogModel, error) {
	boolMustQuery := elastic.NewBoolQuery().Must(this.quertList...)
	rsp, err := AppInit.GetEsClient().Search().Index("bookslogs").
		//Query(urlQuery).  单参数
		Query(boolMustQuery).Size(this.size). // 多参数
		Do(context.Background())
	return MapToLogs(rsp), err
}
