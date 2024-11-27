package main

import (
	"github.com/RE-PIT-BEM/re-pit-backend/cmd/http"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.Load()
	gin.SetMode(config.GetString("gin-mode"))

	http.NewRestHTTPServer()
}
