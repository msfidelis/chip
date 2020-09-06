package healthcheck

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ok godoc
// @Summary Return 200 status Ok in healthcheck
// @Tags Healthcheck
// @Produce json
// @Router /healthcheck [get]
func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

// FaultRandom godoc
// @Summary Inject common errors in healthcheck
// @Tags Healthcheck
// @Produce json
// @Router /healthcheck/fault [get]
// @Tag healthcheck
func FaultRandom(c *gin.Context) {
	status := []int{
		http.StatusServiceUnavailable,
		http.StatusOK,
	}

	n := rand.Int() % len(status)

	c.JSON(status[n], gin.H{
		"status": status[n],
	})
}

// FaultSoft godoc
// @Summary Inject ocasional erros in healthcheck
// @Tags Healthcheck
// @Produce json
// @Router /healthcheck/fault/soft [get]
func FaultSoft(c *gin.Context) {
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

	c.JSON(status[n], gin.H{
		"status": status[n],
	})
}

// Error godoc
// @Summary Return 500 Error Status Code
// @Tags Healthcheck
// @Produce json
// @Router /healthcheck/error [get]
func Error(c *gin.Context) {
	c.JSON(http.StatusServiceUnavailable, gin.H{
		"status": http.StatusServiceUnavailable,
	})
}
