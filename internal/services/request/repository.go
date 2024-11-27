package request

import "github.com/RE-PIT-BEM/re-pit-backend/internal/model/domain"

type RequestRepository interface {
	Create(*domain.Request) error
	FindAll() ([]domain.Request, error)
	FindByID(id string) (domain.Request, error)
	Update(request *domain.Request) error
	Delete(request *domain.Request) error
	FindByUserID(userID int) ([]domain.Request, error)
}
