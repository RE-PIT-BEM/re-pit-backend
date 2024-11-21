package http

import (
	"fmt"
	"net/http"

	"github.com/RE-PIT-BEM/re-pit-backend/internal/config"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/database"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/middleware"

	auth_delivery "github.com/RE-PIT-BEM/re-pit-backend/internal/services/auth/delivery"
	request_delivery "github.com/RE-PIT-BEM/re-pit-backend/internal/services/request/delivery"

	auth_usecase "github.com/RE-PIT-BEM/re-pit-backend/internal/services/auth/usecase"
	request_usecase "github.com/RE-PIT-BEM/re-pit-backend/internal/services/request/usecase"

	request_repository "github.com/RE-PIT-BEM/re-pit-backend/internal/services/request/repository"
	user_repository "github.com/RE-PIT-BEM/re-pit-backend/internal/services/user/repository"
	"github.com/gin-gonic/gin"
)

func NewRestHTTPServer() {
	config := config.Load()
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

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
	// db.AutoMigrate(&domain.Request{})

	// Repository Init
	userRepo := user_repository.NewUserRepository(db)
	requestRepo := request_repository.NewRequestRepository(db)
	// Usecase Init
	authUsecase := auth_usecase.NewAuthUsecase(userRepo)
	requestUsecase := request_usecase.NewRequestUsecase(requestRepo)
	// Delivery Init
	auth_delivery.NewAuthDelivery(router, authUsecase)
	request_delivery.NewRequestDelivery(router, requestUsecase)

	// Run Server
	var httpAddress string
	if config.GetString("env") == "LOCAL" {
		httpAddress = fmt.Sprintf("%s:%s", config.GetString("host"), config.GetString("port"))
	} else {
		httpAddress = fmt.Sprintf(":%s", config.GetString("port"))
	}

	router.Run(httpAddress)
}
