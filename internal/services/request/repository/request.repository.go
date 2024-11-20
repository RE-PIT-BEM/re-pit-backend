package repository

import (
	"github.com/RE-PIT-BEM/re-pit-backend/internal/constant"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/model/domain"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/services/request"
	"gorm.io/gorm"
)

type RequestRepository struct {
	db *gorm.DB
}

func NewRequestRepository(db *gorm.DB) request.RequestRepository {
	return &RequestRepository{db}
}

func (r *RequestRepository) Create(request *domain.Request) error {
	request.RequestStatus = constant.REQUEST_STATUS_PENDING

	err := r.db.Create(&request).Error
	return err
}

func (r *RequestRepository) FindAll() ([]domain.Request, error) {
	var requests []domain.Request
	err := r.db.Find(&requests).Error
	return requests, err
}

func (r *RequestRepository) FindByID(id int) (domain.Request, error) {
	var request domain.Request
	err := r.db.First(&request, id).Error
	return request, err
}

func (r *RequestRepository) Update(request *domain.Request) error {
	err := r.db.Save(&request).Error
	return err
}

func (r *RequestRepository) Delete(request *domain.Request) error {
	err := r.db.Delete(&request).Error
	return err
}

func (r *RequestRepository) FindByUserID(userID int) ([]domain.Request, error) {
	var requests []domain.Request
	err := r.db.Where("user_id = ?", userID).Find(&requests).Error
	return requests, err
}
