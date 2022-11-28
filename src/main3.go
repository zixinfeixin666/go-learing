package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

//文件的上传和返回
func main() {
	router := gin.Default()
	// 给表单限制上传大小 (默认 32 MiB)
	//router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {

		// 单文件
		file, _ := c.FormFile("file")
		//c.SaveUploadedFile(file, "./"+file.Filename)
		//name := c.PostForm("name")
		in, _ := file.Open()
		defer in.Close()
		out, _ := os.Create("./" + file.Filename)
		defer out.Close()
		io.Copy(out, in)

		//后端给前端返回文件内容
		//fmt.Sprintf("attachment;filename=%s",filename)) :对下载的文件重命名
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s", file.Filename))
		c.File("./" + file.Filename)
		//c.JSON(200, gin.H{
		//	"msg:": file,
		//	"name": name,
		//})
	})
	router.Run(":8080")
}
