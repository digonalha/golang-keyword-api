package main

import (
	keyword_controller "keyword-api/src/controllers"

	_ "keyword-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Keyword API
// @version 1.0
// @description API for keyword management
// @host localhost:8001
// @BasePath /
func main() {
	server := gin.Default()
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRoutes := server.Group("/users")
	{
		userRoutes.GET("/:userId/keywords", keyword_controller.GetKeywordsByUserId)
	}

	keywordRoutes := server.Group("/keywords")
	{
		keywordRoutes.POST("", keyword_controller.CreateKeyword)
		keywordRoutes.PUT("/:id", keyword_controller.UpdateKeyword)
		keywordRoutes.DELETE("/:id", keyword_controller.DeleteKeyword)
	}

	server.Run(":8001")

}
