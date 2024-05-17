package reflect

import (
	"github.com/gin-gonic/gin"
)

type Request struct {
	Method  string              `json:"method" binding:"required"`
	Params  string              `json:"params" binding:"required"`
	Headers map[string][]string `json:"headers" binding:"required"`
	// Cookies []*http.Cookie      `json:"cookies" binding:"required"`
	Body string `json:"body" binding:"required"`
	Path string `json:"path" binding:"required"`
}

func Read(c *gin.Context) Request {

	var request Request

	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])

	request.Method = c.Request.Method
	request.Headers = c.Request.Header
	// request.Cookies = c.Request.Cookies()
	request.Path = c.Request.URL.Path
	request.Body = reqBody
	request.Params = c.Request.URL.RawQuery

	return request

}
