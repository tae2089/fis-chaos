package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fis"
	"github.com/aws/aws-sdk-go-v2/service/fis/types"
	"github.com/tae2089/fis-chaos/config"
	"github.com/tae2089/fis-chaos/domain"
	fisTypes "github.com/tae2089/fis-chaos/types"
)

type fisUseCasempl struct {
	fisClient      *fis.Client
	contextTimeout time.Duration
}

func NewFisUseCase(fisClient *fis.Client, contextTimeout time.Duration) domain.FisUsecase {
	return &fisUseCasempl{
		fisClient:      fisClient,
		contextTimeout: contextTimeout,
	}
}

func (f *fisUseCasempl) StartExperiment(ctx context.Context, experimentDto domain.ExperimentDto) error {
	_, err := f.fisClient.StartExperiment(ctx, &fis.StartExperimentInput{
		ExperimentTemplateId: aws.String(experimentDto.TemplateID),
		Tags: map[string]string{
			"Name": experimentDto.Name,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (f *fisUseCasempl) CreateStressChaos(ctx context.Context, stressChaosDto domain.StressChaosDto) (string, error) {
	_, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()
	chaosMeshCfg := config.GetChaosMeshCfg()
	parameters := createChaosMeshParametes(chaosMeshCfg, fisTypes.Stresschaos, stressChaosDto.GetSpec())
	templateID, err := f.createChaosMeshTemplate(stressChaosDto, parameters, chaosMeshCfg.KubernetesARN)
	if err != nil {
		return "", err
	}
	return templateID, nil
}

// CreatePodChaos implements domain.FisUsecase.
func (f *fisUseCasempl) CreatePodChaos(ctx context.Context, podChaos domain.PodChaosDto) (string, error) {
	_, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()
	chaosMeshCfg := config.GetChaosMeshCfg()
	parameters := createChaosMeshParametes(chaosMeshCfg, fisTypes.Stresschaos, podChaos.GetSpec())
	templateID, err := f.createChaosMeshTemplate(podChaos, parameters, chaosMeshCfg.KubernetesARN)
	if err != nil {
		return "", err
	}
	return templateID, nil
}

func (f *fisUseCasempl) createChaosMeshTemplate(chaosReader domain.ChaosReaderable, params domain.ChaosMeshParameters, kubernetesARN string) (string, error) {
	maxDuration, err := convertToTime(chaosReader.GetSec())
	date := time.Now()
	if err != nil {
		return "", err
	}
	output, err := f.fisClient.CreateExperimentTemplate(context.Background(), &fis.CreateExperimentTemplateInput{
		Tags: map[string]string{
			"Name": chaosReader.GetExperimentName(),
		},
		Description: aws.String(chaosReader.GetExperimentName() + " - " + date.Format("2023-05-06")),
		Actions: map[string]types.CreateExperimentTemplateActionInput{
			chaosReader.GetExperimentName(): {
				Description: aws.String(""),
				ActionId:    aws.String(fisTypes.InjectKubernetesCustomResource),
				Parameters: map[string]string{
					"maxDuration":          maxDuration,
					"kubernetesApiVersion": params.APIVersion,
					"kubernetesKind":       params.Kind,
					"kubernetesNamespace":  params.Namespace,
					"kubernetesSpec":       params.Spec,
				},
				Targets: map[string]string{
					"Cluster": "k8s",
				},
			},
		},
		RoleArn: aws.String(params.RoleARN),
		StopConditions: []types.CreateExperimentTemplateStopConditionInput{
			{
				Source: aws.String("none"),
			},
		},
		Targets: map[string]types.CreateExperimentTemplateTargetInput{
			"k8s": {
				ResourceType:  aws.String("aws:eks:cluster"),
				SelectionMode: aws.String("ALL"),
				ResourceArns:  aws.ToStringSlice([]*string{aws.String(kubernetesARN)}),
			},
		},
	})
	if err != nil {
		return "", err
	}
	return *output.ExperimentTemplate.Id, nil
}

func convertToTime(sec int) (string, error) {
	fmt.Println(sec)
	if sec < 60 || sec > 43200 {
		return "", fmt.Errorf("Input is out of range. It should be between 1 and 43200")
	}
	hours := sec / 3600
	minutes := (sec % 3600) / 60
	seconds := sec % 60

	return fmt.Sprintf("PT%02dH%02dM%02dS", hours, minutes, seconds), nil
}

func createChaosMeshParametes(cfg *config.ChaosMeshConfig, kind, spec string) domain.ChaosMeshParameters {
	return domain.ChaosMeshParameters{
		APIVersion: fisTypes.ChaosMeshAPIVersionV1alpha1,
		Kind:       kind,
		Namespace:  cfg.NameSpace,
		Spec:       spec,
		RoleARN:    cfg.RoleARN,
	}
}
