package main

import (
	"BackEnd/loaders"
	r "BackEnd/routers"

	"github.com/gin-gonic/gin"
)

//	@title			bbs-backend API Docs
//	@version		1.0
//	@description	A Bulletin Board System backend base on golang
//  @description	BBS 前後端溝通介面，每個 row 代表一個可使用的 API 實體
//	@schemes		http,https
//	@host			localhost:8080
//	@host			127.0.0.1:8080
//	@BasePath		/api/v1

func main() {
	loaders.Init()

	server := gin.Default()

	// 使用中间件设置CORS头
	server.Use(corsMiddleware())

	r.Routes(server)

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	server.Run(":8080")
}

// 定义CORS中间件
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
