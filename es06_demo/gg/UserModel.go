package gg

import (
	"github.com/graphql-go/graphql"
	"log"
)

type UserModel struct {
	Id   int
	Name string
}

func NewUserModel() *UserModel {
	return &UserModel{Id: 101, Name: "test"}
}

//映射实体  第一步:
func NewUserModelGraphQL() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "UserModel",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.Int},
			"name": &graphql.Field{Type: graphql.String},
		},
	})
}

//创建 查询 对象
func NewUserModelQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "UserQuery",
		Fields: graphql.Fields{
			//取值内容在Resolve里面
			"User": &graphql.Field{Type: NewUserModelGraphQL(), Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				return NewUserModel(), nil
			}},
		},
	})
}

//创建查询规则
func NewUserQuerySchema() graphql.Schema {
	s, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: NewUserModelQuery(),
	})
	if err != nil {
		log.Fatal(err)
	}
	return s
}
