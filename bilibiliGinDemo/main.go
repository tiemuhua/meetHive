package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	engine:=gin.Default()

	engine.LoadHTMLGlob("templete/*")
	
	engine.GET("index", func(context *gin.Context) {
		context.HTML(http.StatusOK,"index.html",gin.H{
			"name":"admin",
			"age":24,
		})
	})

	engine.GET("/helloWorld", func(context *gin.Context) {
		context.String(http.StatusOK,"hello world!")
	})

	engine.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"name":"admin",
			"age":24,
		})
	})

	engine.Run()
}
