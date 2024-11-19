package request

import "github.com/RE-PIT-BEM/re-pit-backend/internal/model/domain"

type RequestRepository interface {
	Create(*domain.Request) error
}
