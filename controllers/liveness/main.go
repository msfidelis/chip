package liveness

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Liveness struct {
	Status string `json:"status" binding:"required"`
}

// Ok godoc
// @Summary Return 200 status Ok in Liveness
// @Tags Liveness
// @Produce json
// @Success 200 {object} Liveness
// @Router /liveness [get]
func Ok(c *gin.Context) {
	var response Liveness
	response.Status = "Live"
	c.JSON(http.StatusOK, response)
}

// Error godoc
// @Summary Return 500 Error Status Code
// @Tags Liveness
// @Produce json
// @Success 200 {object} Liveness
// @Error 503 {object} Liveness
// @Router /liveness/error [get]
func Error(c *gin.Context) {
	var response Liveness
	response.Status = "Dead"
	c.JSON(http.StatusServiceUnavailable, response)
}
