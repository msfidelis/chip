package system

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// // System godoc
// // @Summary Return 500 Error Status Code
// // @Tags System
// // @Produce json
// // @Success 200 {object} system.Capabilities
// // @Router /system [get]
// func System(c *gin.Context) {

// 	var info = system.Info()
// 	// var si sysinfo.SysInfo
// 	// si.GetSysInfo()

// 	c.JSON(http.StatusOK, info)
// 	// c.JSON(http.StatusOK, si)
// }

// Environment godoc
// @Summary Dump all environment variables in container
// @Produce json
// @Tags System
// @Router /system/environment [get]
func Environment(c *gin.Context) {
	c.JSON(http.StatusOK, os.Environ())
}
