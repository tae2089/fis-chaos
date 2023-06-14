package domain

import "context"

type Fis struct {
	Sec            int    `json:"sec" binding:"required"`
	ExperimentName string `json:"experimentName" binding:"required"`
}

type FisUsecase interface {
	CreateStressChaos(ctx context.Context, stressChaos StressChaosDto) (string, error)
	StartExperiment(ctx context.Context, experimentDto ExperimentDto) error
}
