package router

import (
	"github.com/fibbery/go-web/router/handler/sd"
	"github.com/fibbery/go-web/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mdw ...gin.HandlerFunc) {

	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mdw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "not found")
	})

	sdvc := g.Group("/sd")
	{
		sdvc.GET("/health", sd.HealthCheck)
		sdvc.GET("/disk", sd.DiskCheck)
		sdvc.GET("/cpu", sd.CPUCheck)
		sdvc.GET("/ram", sd.RAMCheck)
	}
}
