# fis-chaos

## Description

This repository is a server that creates and executes templates using AWS FIS and Chaos Mesh.

## Usage

Usage in the local environment:

Copy the .env.development file and create a new file called .env.
We determine whether to use the .env file or not based on the environment variable APP_ENV. If you want to use the .env file, simply set the value of APP_ENV to either blank or "local".

Usage in the development, release(production) environment:

If you set environment values such as development, release, or production, they will be configured based on the environment variable values maintained in the operating system.

Run the code.

```go
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


```

## Author

👤 tae2089 tae2089002@gmail.com

Github: @tae2089

## 🤝 Contributing

Contributions, issues and feature requests are welcome!
Feel free to check issues page.
