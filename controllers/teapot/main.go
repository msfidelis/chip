package teapot

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type Teapot struct {
	Body string 
}

// Teapot godoc
// @Summary Return 200 status Teapot in healthcheck
// @Tags ImATeaPot
// @Produce json
// @Success 200 {object} Healthcheck
// @Router /whoami [get]
func Get(c *gin.Context) {

	var response Teapot

	response.Body = "" + 
"             ;,'\n"  +
"     _o_    ;:;'\n" +
" ,-.'---`.__ ;\n" +
"((j`=====',-'\n" +
" `-\\     /\n" +
"    `-=-' " 

	c.String(http.StatusTeapot, response.Body)
}

