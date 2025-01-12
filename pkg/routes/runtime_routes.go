package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

func SetupRuntimeRoutes(router *gin.Engine) {
	runTimeRoutes := router.Group("/runtime")
	{
		runTimeRoutes.GET("/", func(c *gin.Context) {
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)

			c.JSON(http.StatusOK, gin.H{
				"goroutines":             runtime.NumGoroutine(),
				"memory_allocated":       memStats.Alloc,
				"memory_total_allocated": memStats.TotalAlloc,
				"memory_heap_allocated":  memStats.HeapAlloc,
				"memory_heap_system":     memStats.HeapSys,
				"memory_stack_system":    memStats.StackSys,
				"gc_cycles":              memStats.NumGC,
			})
		})
		runTimeRoutes.GET("/debug/pprof/*profile", func(c *gin.Context) {
			http.DefaultServeMux.ServeHTTP(c.Writer, c.Request)
		})
	}
}
