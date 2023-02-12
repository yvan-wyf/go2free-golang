package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wxcloudrun-golang/internal/openai"
)

type ChatRequest struct {
	Question string `question`
}

func main() {
	//if err := db.Init(); err != nil {
	//	panic(fmt.Sprintf("mysql init failed with %+v", err))
	//}

	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	r.POST("/api/chat", func(c *gin.Context) {
		var json ChatRequest

		if err := c.ShouldBindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp := openai.ChatAPI(json.Question)

		if resp.Choices == nil || len(resp.Choices) == 0 {
			c.JSON(http.StatusOK, gin.H{"success": false, "answer": "抱歉，无法回答你的问题"})
		}

		answer := resp.Choices[0].Text
		c.JSON(http.StatusOK, gin.H{"success": true, "answer": answer})
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":80")
}
