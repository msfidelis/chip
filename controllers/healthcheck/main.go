package healthcheck

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Healthcheck struct {
	Status      int    `json:"status" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// Ok godoc
// @Summary Return 200 status Ok in healthcheck
// @Tags Healthcheck
// @Produce json
// @Success 200 {object} Healthcheck
// @Router /healthcheck [get]
func Ok(c *gin.Context) {

	var response Healthcheck
	response.Status = http.StatusOK
	response.Description = "default"

	c.JSON(http.StatusOK, response)
}

// FaultRandom godoc
// @Summary Inject common errors in healthcheck
// @Tags Healthcheck
// @Produce json
// @Router /healthcheck/fault [get]
// @Success 200 {object} Healthcheck
// @Error 503 {object} Healthcheck
// @Tag healthcheck
func FaultRandom(c *gin.Context) {
	var response Healthcheck

	status := []int{
		http.StatusServiceUnavailable,
		http.StatusOK,
	}

	n := rand.Int() % len(status)

	response.Status = status[n]
	response.Description = "fault injection"

	c.JSON(status[n], response)
}

// FaultSoft godoc
// @Summary Inject ocasional erros in healthcheck
// @Tags Healthcheck
// @Produce json
// @Success 200 {object} Healthcheck
// @Error 503 {object} Healthcheck
// @Router /healthcheck/fault/soft [get]
func FaultSoft(c *gin.Context) {
	var response Healthcheck
	status := []int{
		http.StatusServiceUnavailable,
		http.StatusOK,
		http.StatusOK,
		http.StatusOK,
		http.StatusOK,
		http.StatusOK,
		http.StatusOK,
		http.StatusOK,
		http.StatusOK,
		http.StatusOK,
		http.StatusOK,
		http.StatusOK,
	}

	n := rand.Int() % len(status)

	response.Status = status[n]
	response.Description = "fault injection - soft mode with ocasional failures"

	c.JSON(status[n], response)
}

// Error godoc
// @Summary Return 500 Error Status Code
// @Tags Healthcheck
// @Produce json
// @Success 200 {object} Healthcheck
// @Error 503 {object} Healthcheck
// @Router /healthcheck/error [get]
func Error(c *gin.Context) {
	var response Healthcheck
	response.Status = http.StatusServiceUnavailable
	response.Description = "intentional error"
	c.JSON(http.StatusServiceUnavailable, response)
}
