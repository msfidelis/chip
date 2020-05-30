package healthcheck

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

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

func Error(c *gin.Context) {
	c.JSON(http.StatusServiceUnavailable, gin.H{
		"status": http.StatusServiceUnavailable,
	})
}
