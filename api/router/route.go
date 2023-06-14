package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/fis-chaos/config"
)

func SetUp(timeout time.Duration, g *gin.Engine) {

	publicRouter := g.Group("/fis")
	healthRouter := g.Group("/")
	fisClient := config.GetFisClient()
	newFisRouter(fisClient, timeout, publicRouter)
	newHealthRouter(timeout, healthRouter)
}
