package main

import (
	"github.com/gin-gonic/gin"
)

// func main() {

// 	log.Printf("hello world")
// }
// func init() {
// 	file := "./" + "message" + ".txt"
// 	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.SetOutput(logFile) // 将文件设置为log输出的文件
// 	log.SetPrefix("[qSkipTool]")
// 	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
// 	return
// }

// 实现一个以gin为框架的http服务demo
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!\n")
	})
	r.Run("0.0.0.0:8888") // listen and serve on 0.0.0.0:8080
}
