package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func middle() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("我在方法前：1")
		context.Next()
		fmt.Println("我在方法后：1")
	}
}

func middleTwo() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("我在方法前：2")
		context.Next()
		fmt.Println("我在方法后：2")
	}
}

func main() {
	r := gin.Default()
	//测试地址：http://localhost:8080/v1/test
	v1 := r.Group("v1").Use(middle()).Use(middleTwo())
	v1.GET("test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": true,
		})
	})
	r.Run(":8080")
}
