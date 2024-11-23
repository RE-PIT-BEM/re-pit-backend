package delivery

import (
	"fmt"
	"net/http"

	"github.com/RE-PIT-BEM/re-pit-backend/internal/middleware"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/model/dto"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/services/request"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		errorMassages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("Error in field %s, condition: %s", e.Field(), e.ActualTag())
			errorMassages = append(errorMassages, message)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":       "Bad Request!",
			"errorMessages": errorMassages,
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
