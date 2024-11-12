package http

import (
	"net/http"

	"github.com/RE-PIT-BEM/re-pit-backend/internal/config"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/database"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/services/auth/delivery"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/services/auth/usecase"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/services/user/repository"
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

	// Database Init
	db, err := database.NewDatabaseConnection()
	if err != nil {
		panic(err)
	}

	// Migration
	// db.AutoMigrate(&domain.User{})

	// Repository Init
	userRepo := repository.NewUserRepository(db)
	// Usecase Init
	authUsecase := usecase.NewAuthUsecase(userRepo)

	// Delivery Init
	delivery.NewAuthDelivery(router, authUsecase)

	return router
}
