package reflection

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Params  string `json:"params" binding:"required"`
	Headers string `json:"headers" binding:"required"`
}

func Get(c *gin.Context) {

	// var request Request
	// request.Headers = c.Request.Params

	c.JSON(http.StatusOK, c.Request.Header)
}

func Post(c *gin.Context) {

}

func Put(c *gin.Context) {

}

func Patch(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
