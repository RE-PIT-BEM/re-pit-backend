package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/RE-PIT-BEM/re-pit-backend/internal/lib"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/model/dto"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/services/auth"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/services/user"
	"github.com/golang-jwt/jwt/v5"
)

type AuthUsecase struct {
	userRepo user.UserRepository
}

func NewAuthUsecase(userRepo user.UserRepository) auth.AuthUsecase {
	return &AuthUsecase{userRepo}
}

func (u *AuthUsecase) Login(loginRequest *dto.LoginRequestDTO) (dto.LoginResponseDTO, error) {
	// Laod Config
	response := dto.LoginResponseDTO{}

	// Check If User Found
	user, err := u.userRepo.FindByNIM(loginRequest.NIM)
	if err != nil {
		fmt.Println(err.Error())
		return response, errors.New("Invalid NIM or Password!")
	}

	// Check if password is right
	err = lib.CheckPasswordHash(loginRequest.Password, user.Password)
	if err != nil {
		fmt.Println(err.Error())
		return response, errors.New("Invalid NIM or Password!")
	}

	// Creating token
	token, err := lib.CreateJWTToken(jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 10).Unix(),
	})
	if err != nil {
		fmt.Println(err.Error())
		return response, errors.New("Invalid NIM or Password!")
	}

	response.Token = token
	response.User.NIM = user.NIM
	response.User.Name = user.Name
	response.User.Department = user.Department
	response.User.Major = user.Major
	response.User.Role = user.Role

	return response, nil
}

func (u *AuthUsecase) Authorize(userID int) (dto.AuthorizeResponseDTO, error) {
	var response dto.AuthorizeResponseDTO
	userData, err := u.userRepo.FindById(userID)

	response.NIM = userData.NIM
	response.Name = userData.Name
	response.Department = userData.Department
	response.Major = userData.Major
	response.Role = userData.Role

	if err != nil {
		return dto.AuthorizeResponseDTO{}, err
	}

	return response, nil
}
