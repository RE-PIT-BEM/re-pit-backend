package request

import "github.com/RE-PIT-BEM/re-pit-backend/internal/model/dto"

type RequestUsecase interface {
	CreateRequest(req *dto.CreateRequestDTO, userId int) error
}
