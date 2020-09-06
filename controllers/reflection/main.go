package reflection

import (
	"chip/libs/reflect"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get godoc
// @Summary Retun request parameters in response for transformation tests
// @Tags Reflection
// @Produce json
// @Success 200 {object} reflect.Request
// @Router /reflection [get]
func Get(c *gin.Context) {
	request := reflect.Read(c)
	c.JSON(http.StatusOK, request)
}

// Post godoc
// @Summary Retun request parameters in response for transformation tests
// @Tags Reflection
// @Produce json
// @Success 200 {object} reflect.Request
// @Router /reflection [post]
func Post(c *gin.Context) {
	request := reflect.Read(c)
	c.JSON(http.StatusOK, request)
}

// Put godoc
// @Summary Retun request parameters in response for transformation tests
// @Tags Reflection
// @Produce json
// @Success 200 {object} reflect.Request
// @Router /reflection [put]
func Put(c *gin.Context) {
	request := reflect.Read(c)
	c.JSON(http.StatusOK, request)
}

// Patch godoc
// @Summary Retun request parameters in response for transformation tests
// @Tags Reflection
// @Produce json
// @Success 200 {object} reflect.Request
// @Router /reflection [patch]
func Patch(c *gin.Context) {
	request := reflect.Read(c)
	c.JSON(http.StatusOK, request)
}

// Delete godoc
// @Summary Retun request parameters in response for transformation tests
// @Tags Reflection
// @Produce json
// @Success 200 {object} reflect.Request
// @Router /reflection [delete]
func Delete(c *gin.Context) {
	request := reflect.Read(c)
	c.JSON(http.StatusOK, request)
}
