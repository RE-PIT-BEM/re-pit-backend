package http

import (
	"net/http"

	"github.com/RE-PIT-BEM/re-pit-backend/internal/config"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/database"
	"github.com/gin-gonic/gin"
)

func NewRestHTTPServer() *gin.Engine {
	config.Load()
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello From RE-PIT Backend!!",
		})
	})

	_, err := database.NewDatabaseConnection()
	if err != nil {
		panic(err)
	}

	// Migration
	// db.AutoMigrate(&domain.User{})

	return router
}
