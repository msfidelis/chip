package burn

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// Cpu godoc
// @Summary Burn CPU for Loadtests and Auto Scaling Tests
// @Produce json
// @Tags Loadtest
// @Router /burn/cpu [get]
func Cpu(c *gin.Context) {

	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)

	quit := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for {
				select {
				case <-quit:
					return
				default:
				}
			}
		}()
	}

	time.Sleep(3 * time.Second)
	for i := 0; i < n; i++ {
		quit <- true
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "On Fire",
	})

}

// Cpu godoc
// @Summary Burn RAM for Loadtests and Auto Scaling Tests
// @Produce json
// @Tags Loadtest
// @Router /burn/ram [get]
func Mem(c *gin.Context) {

	var s []int

	sum := 1
	for sum < 9999998 {
		sum += 1
		s = append(s, sum)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "On Fire",
	})

}
