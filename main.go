package main

import (
	keyword_handler "keyword-api/src/handlers"

	"github.com/gin-gonic/gin"
)

func createRouter() {
	router := gin.Default()

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/:userId/keywords", keyword_handler.GetKeywordsByUserId)
	}

	keywordRoutes := router.Group("/keywords")
	{
		keywordRoutes.POST("", keyword_handler.CreateKeyword)
		keywordRoutes.PUT("/:id", keyword_handler.UpdateKeyword)
		keywordRoutes.DELETE("/:id", keyword_handler.DeleteKeyword)
	}

	router.Run(":8001")
}

func main() {
	createRouter()
}
