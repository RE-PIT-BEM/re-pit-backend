package auth

import "github.com/RE-PIT-BEM/re-pit-backend/internal/model/dto"

type AuthUsecase interface {
	Login(loginRequest *dto.LoginRequestDTO) (string, error)
}
