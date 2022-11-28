package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//测试地址：http://localhost:8080/v1/test
	v1 := r.Group("v1")
	v1.GET("test", func(context *gin.Context) {
		fmt.Println("我在分组方法内")
		context.JSON(200, gin.H{
			"message": true,
		})
	})
	r.Run(":8080")
}
