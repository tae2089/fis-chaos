package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/fis-chaos/domain"
)

type FisHandler struct {
	FisUsecase domain.FisUsecase
}

func (fh *FisHandler) CreateStressChaosTemplate(c *gin.Context) {

	var stressChaosDto domain.StressChaosDto

	if err := c.ShouldBindJSON(&stressChaosDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := fh.FisUsecase.CreateStressChaos(c, stressChaosDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"isSuccess": true, "chaosTemplateID": id})
}

func (fh *FisHandler) StartExperiment(c *gin.Context) {
	var experimentDto domain.ExperimentDto

	if err := c.ShouldBindJSON(&experimentDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := fh.FisUsecase.StartExperiment(c, experimentDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"isSuccess": true})
}
