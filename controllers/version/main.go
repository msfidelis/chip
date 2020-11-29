package version

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get godoc
// @Summary Return version for container
// @Produce json
// @Tags Version
// @Router /version [get]
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": "v2",
	})
}
