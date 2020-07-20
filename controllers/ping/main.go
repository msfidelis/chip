package ping

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Tcp(c *gin.Context) {

	host := c.Param("host")
	port := c.Param("port")
	connection_string := net.JoinHostPort(host, port)

	fmt.Println("Testing connection on: " + connection_string)

	conn, err := net.DialTimeout("tcp", connection_string, 3*time.Second)
	defer conn.Close()

	if err != nil {
		fmt.Println("Error in connection:", err)

		c.JSON(http.StatusOK, gin.H{
			"host":     host,
			"port":     port,
			"protocol": "tcp",
			"status":   "error to connect",
		})

	} else {
		fmt.Println("Success to connect: " + connection_string)
		c.JSON(http.StatusOK, gin.H{
			"host":     host,
			"port":     port,
			"protocol": "tcp",
			"status":   "connected",
		})
	}

}
