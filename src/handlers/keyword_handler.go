package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateKeyword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "keyword created"})
}
