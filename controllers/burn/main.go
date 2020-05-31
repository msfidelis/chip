package burn

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

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

	time.Sleep(2 * time.Second)
	for i := 0; i < n; i++ {
		quit <- true
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "On Fire",
	})

}
