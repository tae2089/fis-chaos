package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/fis-chaos/config"
)

type HealthHandler struct{}

func (fh *HealthHandler) CheckHealth(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"isLive": true})
}

func (fg *HealthHandler) CheckEnv(c *gin.Context) {
	cfg := config.GetChaosMeshCfg()
	c.JSON(http.StatusOK, gin.H{"envs": cfg})
}
