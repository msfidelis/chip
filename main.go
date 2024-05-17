package main

import (
	"chip/controllers/burn"
	"chip/controllers/healthcheck"
	"chip/controllers/liveness"
	"chip/controllers/logging"
	"chip/controllers/ping"
	"chip/controllers/proxy"
	"chip/controllers/readiness"
	"chip/controllers/reflection"
	"chip/controllers/system"

	// "chip/controllers/system"
	"chip/controllers/teapot"
	"chip/controllers/version"
	loggerInternal "chip/libs/logger"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Depado/ginprom"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

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

	router := gin.New()

	// Logger
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	// Memory Cache Singleton
	c := memory_cache.GetInstance()

	// Readiness Probe Mock Config
	probe_time_raw := os.Getenv("READINESS_PROBE_MOCK_TIME_IN_SECONDS")
	if probe_time_raw == "" {
		probe_time_raw = "5"
	}
	probe_time, err := strconv.ParseUint(probe_time_raw, 10, 64)
	if err != nil {
		fmt.Println("Environment variable READINESS_PROBE_MOCK_TIME_IN_SECONDS conversion error", err)
	}
	c.Set("readiness.ok", "false", time.Duration(probe_time)*time.Second)

	// Prometheus Exporter Config
	p := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)

	// Logger
	logInternal := loggerInternal.Instance()

	//Middlewares
	router.Use(p.Instrument())
	router.Use(gin.Recovery())
	router.Use(chaos.Load())
	router.Use(logger.SetLogger(
		logger.WithSkipPath([]string{"/skip"}),
		logger.WithUTC(true),
		// logger.WithLogger(func(c *gin.Context, out io.Writer, latency time.Duration) zerolog.Logger {
		// 	return zerolog.New(out).With().
		// 		Str("method", c.Request.Method).
		// 		Str("path", c.Request.URL.Path).
		// 		Dur("latency", latency).
		// 		Logger()
		// }),
	))

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

	// // System
	router.GET("/system", system.System)
	router.GET("/system/environment", system.Environment)

	// Stress Test
	router.GET("/burn/cpu", burn.Cpu)
	router.GET("/burn/ram", burn.Mem)

	// Reflection - DEPRECATED
	router.GET("/reflection", reflection.Get)
	router.POST("/reflection", reflection.Post)
	router.PUT("/reflection", reflection.Put)
	router.PATCH("/reflection", reflection.Patch)
	router.DELETE("/reflection", reflection.Delete)

	// Echo
	router.GET("/echo", reflection.Get)
	router.POST("/echo", reflection.Post)
	router.PUT("/echo", reflection.Put)
	router.PATCH("/echo", reflection.Patch)
	router.DELETE("/echo", reflection.Delete)

	// Logging
	router.GET("/logging", logging.Get)

	//Ping
	router.GET("/ping/:host/:port", ping.Tcp)

	// I'am a Teapot
	router.GET("/whoami", teapot.Get)

	// Proxy
	router.POST("/proxy", proxy.Post)

	// Graceful Shutdown Config
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logInternal.
				Error().
				Str("Error", err.Error()).
				Msg("Failed to listen")
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logInternal.
		Warn().
		Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logInternal.
			Error().
			Str("Error", err.Error()).
			Msg("Server forced to shutdown: ")
	}

	fmt.Println("Server exiting")

}
