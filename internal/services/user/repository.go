package user

import "github.com/RE-PIT-BEM/re-pit-backend/internal/model/domain"

type UserRepository interface {
	FindByNIM(NIM string) (domain.User, error)
}
