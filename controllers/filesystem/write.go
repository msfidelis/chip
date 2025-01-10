package filesystem

import (
	"encoding/base64"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type RequestWrite struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

func WriteFile(c *gin.Context) {
	var request RequestWrite

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	decodedContent, err := base64.StdEncoding.DecodeString(request.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid base64 content"})
		return
	}

	err = os.WriteFile(request.Path, decodedContent, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "file written", "path": request.Path})
}
