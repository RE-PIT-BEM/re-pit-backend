package usecase

import (
	"fmt"

	"github.com/RE-PIT-BEM/re-pit-backend/internal/model/dto"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/services/auth"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/services/user"
)

type AuthUsecase struct {
	userRepo user.UserRepository
}

func NewAuthUsecase(userRepo user.UserRepository) auth.AuthUsecase {
	return &AuthUsecase{userRepo}
}

func (u *AuthUsecase) Login(loginRequest *dto.LoginRequestDTO) (string, error) {
	// Logic Of Login User
	user, err := u.userRepo.FindByNIM(loginRequest.NIM)
	fmt.Println(user)
	return "Ketemu", err
}
