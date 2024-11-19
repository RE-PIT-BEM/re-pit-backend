package delivery

import (
	"net/http"

	"github.com/RE-PIT-BEM/re-pit-backend/internal/middleware"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/model/dto"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/services/request"
	"github.com/gin-gonic/gin"
)

type RequestDelivery struct {
	router  *gin.Engine
	usecase request.RequestUsecase
}

func NewRequestDelivery(router *gin.Engine, usecase request.RequestUsecase) {
	handler := &RequestDelivery{
		router:  router,
		usecase: usecase,
	}

	router.POST("/request", middleware.RequireAuth, handler.CreateRequestHandler)
}

func (r *RequestDelivery) CreateRequestHandler(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")

	var createRequestDTO dto.CreateRequestDTO
	err := ctx.Bind(&createRequestDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request!",
		})
		return
	}

	err = r.usecase.CreateRequest(&createRequestDTO, int(userId.(float64)))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":      "Bad Request!",
			"errorMessage": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Request Created!",
	})
}
