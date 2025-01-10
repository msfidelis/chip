package filesystem

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type RequestDelete struct {
	Path string `json:"path"`
}

func DeleteFile(c *gin.Context) {
	var request RequestDelete
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := os.Remove(request.Path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "file deleted"})
}
