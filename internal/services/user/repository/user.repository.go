package repository

import (
	"github.com/RE-PIT-BEM/re-pit-backend/internal/model/domain"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/services/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) FindByNIM(NIM string) (domain.User, error) {
	var user domain.User
	err := u.db.Where("nim = ?", NIM).First(&user).Error
	return user, err
}
