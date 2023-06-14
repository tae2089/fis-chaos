package router

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/fis"
	"github.com/gin-gonic/gin"
	"github.com/tae2089/fis-chaos/api/handler"
	"github.com/tae2089/fis-chaos/usecase"
)

func newFisRouter(fisClient *fis.Client, timeout time.Duration, group *gin.RouterGroup) {
	fisUseCase := usecase.NewFisUseCase(fisClient, timeout)
	handler := &handler.FisHandler{
		FisUsecase: fisUseCase,
	}
	group.POST("/stress", handler.CreateStressChaosTemplate)
	group.POST("/start", handler.StartExperiment)
}
