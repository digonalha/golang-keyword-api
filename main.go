package main

import (
	"keyword-api/src/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/keyword", handlers.CreateKeyword)

	router.Run(":8001")
}
