package delivery

import (
	"net/http"

	"github.com/RE-PIT-BEM/re-pit-backend/internal/middleware"
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
	router.GET("/authorize", middleware.RequireAuth, handler.AuthorizeHandler)
}

func (d *AuthDelivery) LoginHandler(ctx *gin.Context) {
	var loginReq dto.LoginRequestDTO

	err := ctx.Bind(&loginReq)
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

func (d *AuthDelivery) AuthorizeHandler(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")

	userData, err := d.usecase.Authorize(int(userId.(float64)))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Succesfully Authorized!",
		"data": gin.H{
			"user": userData,
		},
	})
}
