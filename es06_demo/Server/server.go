package main

import (
	"Go_ElasticSearch7/es06_demo/Funs"
	"Go_ElasticSearch7/es06_demo/Middleware"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Use(Middleware.LogMiddleware(), gin.Recovery()) //设置日志

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("mygte", func(fl validator.FieldLevel) bool {
			param := fl.Param()                        //获取参数  mygte=BookPrice1Start 的 BookPrice1Start
			v := fl.Parent().Elem().FieldByName(param) //判断=BookPrice1Start是否合法
			if !v.IsValid() {
				return false
			}
			if fl.Field().Float() >= v.Float() {
				return true
			}
			////fmt.Println(fl.Field().Float(),v.Float())  //获取字段值
			//return false
			return false
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	g := router.Group("/books")
	{
		//图书列表
		g.Handle("GET", "", Funs.LoadBook)

		//http://localhost:8080/books/press/湘潭大学出版社
		g.Handle("GET", "/press/:press", Funs.LoadBookByPress)

		//http://localhost:8080/books/presses/湘潭大学出版社,人民邮电出版社
		g.Handle("GET", "/presses/:press", Funs.LoadBooksByPress)

		//搜索api
		g.Handle("POST", "/search", Funs.SeachBook)
	}

	//查询日志里面的aggs
	loggroup := router.Group("/log/aggs")
	{
		//聚合
		//localhost:8080/log/aggs/max/duration  延迟时间最长的值
		loggroup.Handle("GET", "/:type/:field", Funs.LogAgg)
	}

	//helper
	helper := router.Group("/helper")
	{
		helper.Handle("GET", "/press", Funs.PressList)
	}

	router.StaticFS("/ui", http.Dir("./es06_demo/htmls"))

	router.Run(":8080")
}
