package environment

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, os.Environ())
}
