package main

import (
	"Go_ElasticSearch7/es06_demo/gg"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func main() {

	//queryString := `
	//				query{
	//					User(id:2){
	//						id
	//						name
	//					}
	//				}
	//				`

	//param := graphql.Params{Schema: gg.NewUserQuerySchema(), RequestString: queryString}
	//ret := graphql.Do(param)
	//if ret.HasErrors() {
	//	log.Fatal("data error",ret.Errors)
	//} else {
	//	fmt.Println(ret.Data)
	//}

	//schema := gg.NewUserQuerySchema()   //获取   /users
	schema := gg.NewLogQuerySchema()
	h := handler.New(&handler.Config{
		Schema: &schema,
	})
	/*
		query{
			User(id:1) {
			id
			name
		}
		}
	*/
	router := gin.Default()
	//g := router.Group("/users")
	g := router.Group("/logs")
	{
		g.Handle("POST", "/", func(context *gin.Context) {
			h.ServeHTTP(context.Writer, context.Request)
		})
	}
	router.Run(":8080")
}
