package main

import (
	"chip/controllers/healthcheck"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/healthcheck", healthcheck.Get)
	router.Run()
}
