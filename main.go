package main

import (
	"chip/controllers/burn"
	"chip/controllers/healthcheck"
	"chip/controllers/liveness"
	"chip/controllers/ping"
	"chip/controllers/readiness"
	"chip/controllers/reflection"
	"chip/controllers/system"
	"chip/controllers/version"
	"time"

	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"

	// "github.com/patrickmn/go-cache"
	"chip/libs/memory_cache"
	chaos "github.com/msfidelis/gin-chaos-monkey"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "chip/docs"
)

// @title Chip
// @version 1.0
// @description Cloud Native Toolset Running in Containers.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email matheus@nanoshots.com.br

// @license.name MIT
// @license.url https://github.com/mfidelis/chip/blob/master/LICENSE

// @BasePath /
func main() {

	router := gin.Default()

	// Memory Cache Singleton
	c := memory_cache.GetInstance()

	// Readiness Probe Mock Config

	probe_time := os.GetEnv("READINESS_PROBE_MOCK_TIME")

	c.Set("readiness.ok", "false", 30*time.Second)

	p := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)

	//Middlewares
	router.Use(p.Instrument())
	router.Use(gin.Recovery())
	router.Use(chaos.Load())

	//Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Healthcheck
	router.GET("/healthcheck", healthcheck.Ok)
	router.GET("/healthcheck/fault", healthcheck.FaultRandom)
	router.GET("/healthcheck/fault/soft", healthcheck.FaultSoft)
	router.GET("/healthcheck/error", healthcheck.Error)

	// Liveness
	router.GET("/liveness", liveness.Ok)
	router.GET("/liveness/error", liveness.Error)

	// Readinesscurl
	router.GET("/readiness", readiness.Ok)
	router.GET("/readiness/error", readiness.Error)

	// Version
	router.GET("/version", version.Get)

	// System
	router.GET("/system", system.System)
	router.GET("/system/environment", system.Environment)

	// Stress Test
	router.GET("/burn/cpu", burn.Cpu)
	router.GET("/burn/ram", burn.Mem)

	// Reflection
	router.GET("/reflection", reflection.Get)
	router.POST("/reflection", reflection.Post)
	router.PUT("/reflection", reflection.Put)
	router.PATCH("/reflection", reflection.Patch)
	router.DELETE("/reflection", reflection.Delete)

	//Ping
	router.GET("/ping/:host/:port", ping.Tcp)

	router.Run()
}
