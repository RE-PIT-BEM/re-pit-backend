package delivery

import (
	"net/http"

	"github.com/RE-PIT-BEM/re-pit-backend/internal/model/dto"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/services/auth"
	"github.com/gin-gonic/gin"
)

type AuthDelivery struct {
	router  *gin.Engine
	usecase auth.AuthUsecase
}

func NewAuthDelivery(router *gin.Engine, usecase auth.AuthUsecase) {
	handler := &AuthDelivery{router, usecase}

	router.POST("/login", handler.LoginHandler)
}

func (d *AuthDelivery) LoginHandler(ctx *gin.Context) {
	var loginReq dto.LoginRequestDTO

	err := ctx.BindJSON(&loginReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request!",
		})
		return
	}

	data, err := d.usecase.Login(&loginReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":      "Bad Request!",
			"errorMessage": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Succesfully Login!",
		"data":    data,
	})
}
