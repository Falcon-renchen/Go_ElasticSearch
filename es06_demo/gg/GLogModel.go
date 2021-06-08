package gg

import (
	"github.com/graphql-go/graphql"
	"log"
)

type LogModel struct {
	Ip       string `json:"ip"`
	Status   string `json:"status"`
	Duration string `json:"duration"`
	Method   string `json:"method"`
	Url      string `json:"url"`
	Time     string `json:"time"`
	Agent    string `json:"agent"`
	Referer  string `json:"referer"`
}

func NewLogModelGraphQL() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "LogModel",
		Fields: graphql.Fields{
			"Ip":       StringField(),
			"Status":   StringField(),
			"Duration": StringField(),
			"Method":   StringField(),
			"Url":      StringField(),
			"Time":     StringField(),
			"Msg":      StringField(),
			"Agent":    StringField(),
		},
	})
}

func NewLogModelQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "LogQuery",
		//列表显示
		Fields: graphql.Fields{
			/*
				query{
				    Search {
				        Ip
				        Url
				    }
				}
			*/
			"Searchs": &graphql.Field{Type: graphql.NewList(NewLogModelGraphQL()),
				//不传参数不用写这个
				//Args:StringArg("url"), 单参数
				Args: StringArgs("url", "method"), //多参数
				//Args: graphql.FieldConfigArgument{"id": &graphql.ArgumentConfig{Type: graphql.Int}},
				//和User(id:1)对应
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					//return NewLogService().Search(p.Args["url"].(string)) //单参数
					return NewLogService().Searchs(p.Args["url"], p.Args["method"]) //多参数
				}},
		},
	})
}

//创建查询规则
func NewLogQuerySchema() graphql.Schema {
	s, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: NewLogModelQuery(),
	})
	if err != nil {
		log.Fatal(err)
	}
	return s
}
