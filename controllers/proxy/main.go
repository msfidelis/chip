package proxy

import (
	"chip/libs/httpclient"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Method  string `json:"method"`
	Host    string `json:"host"`
	Path    string `json:"path"`
	Headers []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"headers"`
	Body string `json:"body"`
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Body       string      `json:"body"`
	Headers    http.Header `json:"headers"`
}

// Proxy godoc
// @Summary Proxy Request
// @Tags Proxy
// @Produce json
// @Param message body Request true "Proxy Information"
// @Success 200 {object} Response
// @Router /proxy [post]
func Post(c *gin.Context) {

	var request Request
	var response Response

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	headers := make(map[string][]string)
	for _, header := range request.Headers {
		headers[header.Name] = append(headers[header.Name], header.Value)
	}

	res, body := httpclient.Request(request.Method, request.Host, request.Path, headers, request.Body)

	response.StatusCode = res.StatusCode
	response.Body = body
	response.Headers = res.Header

	c.JSON(http.StatusOK, response)
}
