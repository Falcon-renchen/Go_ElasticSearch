package main

import (
	"Go_ElasticSearch7/es06_demo/gg"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
)

func main() {

	queryString := `
					query{
						User(id:1){
							id
							name
						}
					}
					`
	param := graphql.Params{Schema: gg.NewUserQuerySchema(), RequestString: queryString}
	ret := graphql.Do(param)
	if ret.HasErrors() {
		log.Fatal(ret.Errors)
	} else {
		fmt.Println(ret.Data)
	}
}
