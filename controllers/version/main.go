package version

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Get godoc
// @Summary Return version for container
// @Produce json
// @Tags Version
// @Router /version [get]
func Get(c *gin.Context) {

	version := os.Getenv("VERSION")
	if version == "" {
		version = "v2"
	}

	c.JSON(http.StatusOK, gin.H{
		"version": version,
	})
}
