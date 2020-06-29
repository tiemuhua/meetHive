package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Name string `form:"name" binding:"required,len=6"`
	Age  int    `form:"age" binding:"numeric,min=18,max=100"`
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("./templete/*")
	engine.Static("upload", "./static")

	engine.GET("/upload", func(context *gin.Context) {
		context.HTML(http.StatusOK, "upload.html", nil)
	})

	engine.POST("/upload", func(context *gin.Context) {
		f, err := context.FormFile("file")
		if err != nil {
			log.Println()
		}
		err = context.SaveUploadedFile(f, fmt.Sprintf("uploads/%s", f.Filename))
		fmt.Println(f.Filename + "------------")
		if err != nil {
			context.String(http.StatusOK, "upload file failed->%v", err.Error())
		} else {
			context.String(http.StatusOK, "upload success")
		}

	})
	engine.Run()
}
