package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/zcalusic/sysinfo"

	"chip/libs/system"
)

func Get(c *gin.Context) {

	var info = system.Info()
	// var si sysinfo.SysInfo
	// si.GetSysInfo()

	c.JSON(http.StatusOK, info)
	// c.JSON(http.StatusOK, si)
}