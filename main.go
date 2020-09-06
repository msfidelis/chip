package main

import (
	"chip/controllers/burn"
	"chip/controllers/environment"
	"chip/controllers/healthcheck"
	"chip/controllers/ping"
	"chip/controllers/reflection"
	"chip/controllers/system"
	"chip/controllers/version"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// Healthcheck
	router.GET("/healthcheck", healthcheck.Ok)
	router.GET("/healthcheck/fault", healthcheck.FaultRandom)
	router.GET("/healthcheck/fault/soft", healthcheck.FaultSoft)
	router.GET("/healthcheck/error", healthcheck.Error)

	// Version
	router.GET("/version", version.Get)

	// Environment
	router.GET("/environment", environment.Get)

	// System
	router.GET("/system", system.Get)

	// Stress Test
	router.GET("/burn/cpu", burn.Cpu)
	router.GET("/burn/ram", burn.Mem)

	// Reflection
	router.Any("/reflection", reflection.Get)

	//Ping
	router.GET("/ping/:host/:port", ping.Tcp)

	router.Run()
}
