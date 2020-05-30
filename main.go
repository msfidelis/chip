package main

import (
	"chip/controllers/healthcheck"
	"chip/controllers/system"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	// Healthcheck
	router.GET("/healthcheck", healthcheck.Ok)
	router.GET("/healthcheck/fault", healthcheck.FaultRandom)
	router.GET("/healthcheck/error", healthcheck.Error)

	// System
	router.GET("/system", system.Get)

	router.Run()
}
