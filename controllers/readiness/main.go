package readiness

import (
	"chip/libs/memory_cache"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Readiness struct {
	Status string `json:"status" binding:"required"`
}

// Ok godoc
// @Summary Return 200 status Ok in Liveness
// @Tags Readiness
// @Produce json
// @Success 200 {object} Liveness
// @Router /readiness [get]
func Ok(c *gin.Context) {

	m := memory_cache.GetInstance()
	var response Readiness

	foo, found := m.Get("readiness.ok")
	if found {
		fmt.Println("Não esta pronto")
		fmt.Println(foo)
	} else {
		fmt.Println("Probe ok")
	}

	response.Status = "Ready"
	c.JSON(http.StatusOK, response)
}

// Error godoc
// @Summary Return 500 Error Status Code
// @Tags Readiness
// @Produce json
// @Success 200 {object} Readiness
// @Error 503 {object} Readiness
// @Router /readiness/error [get]
func Error(c *gin.Context) {
	var response Readiness
	response.Status = "Not Ready"
	c.JSON(http.StatusServiceUnavailable, response)
}
