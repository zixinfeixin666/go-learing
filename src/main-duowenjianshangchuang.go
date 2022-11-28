package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.POST("/testUpload", func(context *gin.Context) {
		form, _ := context.MultipartForm()
		file := form.File["file"]
		fmt.Println(file)
		for _, f := range file {
			log.Println(f.Filename)
		}
	})
	r.Run(":8080")

}
