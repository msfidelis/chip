package filesystem

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type RequestCat struct {
	Path string `json:"path"`
}

func Cat(c *gin.Context) {
	var request RequestCat

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content, err := os.ReadFile(request.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, string(content))
}
