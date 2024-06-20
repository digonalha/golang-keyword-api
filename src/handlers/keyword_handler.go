package keyword_handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateKeyword(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "keyword created"})
}

func GetKeywordsByUserId(c *gin.Context) {
	userId := c.Param("userId")
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("List keywords from user %s", userId)})
}

func UpdateKeyword(c *gin.Context) {
	keywordId := c.Param("id")
	c.JSON(http.StatusAccepted, gin.H{"message": fmt.Sprintf("keyword %s updated", keywordId)})
}

func DeleteKeyword(c *gin.Context) {
	keywordId := c.Param("id")
	c.JSON(http.StatusAccepted, gin.H{"message": fmt.Sprintf("keyword %s deleted", keywordId)})
}
