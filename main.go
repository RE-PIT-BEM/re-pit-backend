package main

import (
	"fmt"

	"github.com/RE-PIT-BEM/re-pit-backend/cmd/http"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.Load()
	gin.SetMode(config.GetString("gin-mode"))

	httpRouter := http.NewRestHTTPServer()
	var httpAddress string
	if config.GetString("env") == "LOCAL" {
		httpAddress = fmt.Sprintf("%s:%s", config.GetString("host"), config.GetString("port"))
	} else {
		httpAddress = fmt.Sprintf(":%s", config.GetString("port"))
	}

	httpRouter.Run(httpAddress)
}
