package Middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func LogMiddleware() gin.HandlerFunc {
	//固定代码
	logger := logrus.New()
	logfile, err := os.OpenFile("./es06_demo/mylog", os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		logrus.Fatal(err)
	}
	logger.AddHook(NewEsHook()) //添加钩子
	logger.Out = logfile

	return func(ctx *gin.Context) {
		startTime := time.Now()             //开始时间
		ctx.Next()                          //具体的业务流程
		endTime := time.Now()               //结束时间
		execTime := endTime.Sub(startTime)  //两边一减 = 响应时间
		requestMethod := ctx.Request.Method // 下面是获取
		requestURI := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		requestIP := ctx.ClientIP()
		////日志格式
		//logger.Infof("| %2d | %12v | %14s | %s | %s |",
		//	statusCode,
		//	execTime,
		//	requestIP,
		//	requestMethod,
		//	requestURI,
		//)
		//进行设置字段
		logger.WithField("ip", requestIP).
			WithField("status", statusCode).
			WithField("duration", execTime.Milliseconds()).
			WithField("method", requestMethod).
			WithField("url", requestURI).Info() //hook勾住了url，就会在terminal显示  /books?a=2   类似url

	}
}
