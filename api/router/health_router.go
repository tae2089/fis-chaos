package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/fis-chaos/api/handler"
)

func newHealthRouter(timeout time.Duration, group *gin.RouterGroup) {
	healthRouter := handler.HealthHandler{}
	group.GET("/healthz", healthRouter.CheckHealth)
	group.GET("/envs", healthRouter.CheckEnv)
}
