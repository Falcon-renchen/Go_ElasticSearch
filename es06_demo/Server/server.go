package main

import (
	"Go_ElasticSearch7/es06_demo/Funs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("mygte", func(fl validator.FieldLevel) bool {
			param := fl.Param()
			v := fl.Parent().Elem().FieldByName(param)
			if !v.IsValid() {
				return false
			}
			if fl.Field().Float() >= v.Float() {
				return true
			}
			return false
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	//if v, ok := binding.Validator.Engine().(*validator.Validate);ok {
	//	err := v.RegisterValidation("mygte", func(fl validator.FieldLevel) bool {
	//		param := fl.Param()   //获取参数  mygte=BookPrice1Start 的 BookPrice1Start
	//		v := fl.Parent().Elem().FieldByName(param)  //判断=BookPrice1Start是否合法
	//		if !v.IsValid() {
	//			return false
	//		}
	//		if fl.Field().Float()>=v.Float() {
	//			return true
	//		}
	//		////fmt.Println(fl.Field().Float(),v.Float())  //获取字段值
	//		//return false
	//		return false
	//	})
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

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

	//helper
	helper := router.Group("/helper")
	{
		helper.Handle("GET", "/press", Funs.PressList)
	}

	router.StaticFS("/ui", http.Dir("./htmls"))

	router.Run(":8080")
}
