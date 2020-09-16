package system

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"chip/libs/logging"
	"chip/libs/system"
)

// System godoc
// @Summary Return 500 Error Status Code
// @Tags System
// @Produce json
// @Success 200 {object} system.Capabilities
// @Router /system [get]
func System(c *gin.Context) {
	logging.Debug("message", "getting things kkkk")
	c.JSON(http.StatusOK, system.Info())
}

// Environment godoc
// @Summary Dump all environment variables in container
// @Produce json
// @Tags System
// @Router /system/environment [get]
func Environment(c *gin.Context) {
	c.JSON(http.StatusOK, os.Environ())
}
