package main

import (
	"bbs_backend/loaders"
	r "bbs_backend/routers"

	_ "bbs_backend/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	r.Routes(server)

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	server.Run(":8080")
}
