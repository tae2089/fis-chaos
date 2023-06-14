package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/fis"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type ChaosMeshConfig struct {
	NameSpace     string `env:"CHAOS_MESH_NAMESPACE,required"`
	KubernetesARN string `env:"CHAOS_MESH_KUBERNETES_ARN,required"`
	RoleARN       string `env:"CHAOS_MESH_ROLE_ARN,required"`
}

var chaosMeshCfg = &ChaosMeshConfig{}

func LoadEnv() {

	var err error

	if os.Getenv("APP_ENV") == "local" || os.Getenv("APP_ENV") == "" {
		projectDir := getProjectDir() + "/.env"
		err = godotenv.Load(projectDir)
	}

	if err != nil {
		log.Println("Error loading.env file")
		return
	}

	if err := env.Parse(chaosMeshCfg); err != nil {
		fmt.Printf("err: %+v\n", err)
	}
}

func getProjectDir() string {
	projectDir := ""
	_, filename, _, _ := runtime.Caller(0)
	projectDir = path.Join(path.Dir(filename), "..")
	return projectDir
}

func GetFisClient() *fis.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Panic(err)
	}
	fisClient := fis.NewFromConfig(cfg)
	return fisClient
}

func GetChaosMeshCfg() *ChaosMeshConfig {
	return chaosMeshCfg
}
