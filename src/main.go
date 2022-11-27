package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default() //携带基础中间件启动
	r.GET("/path/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		user := c.DefaultQuery("user", "xk")
		pwd := c.Query("pwd")
		c.JSON(200, gin.H{
			"id":      id,
			"user":    user,
			"pwd":     pwd,
			"success": true,
		})
	})
	r.POST("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "xk")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})
	r.DELETE("/path/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		c.JSON(200, gin.H{
			"id":      id,
			"success": true,
		})
	})
	r.PUT("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "xk")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"user":    user,
			"pwd":     pwd,
			"success": true,
		})
	})
	r.Run(":1010") // listen and serve on 0.0.0.0:8080
}
