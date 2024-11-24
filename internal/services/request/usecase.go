package request

import (
	"github.com/RE-PIT-BEM/re-pit-backend/internal/model/domain"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/model/dto"
	"github.com/gin-gonic/gin"
)

type RequestUsecase interface {
	CreateRequest(req *dto.CreateRequestDTO, userId int) error
	GetAllRequest() ([]domain.Request, error)
	GetRequestByUserID(userID int) ([]domain.Request, error)
	GetRequestByID(ctx *gin.Context, id string) (domain.Request, error)
}
