package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/fis-chaos/api/router"
	"github.com/tae2089/fis-chaos/config"
)

func main() {
	config.LoadEnv()
	timeout := time.Duration(3) * time.Second
	gin := gin.Default()
	router.SetUp(timeout, gin)
	gin.Run(":8080")
}
