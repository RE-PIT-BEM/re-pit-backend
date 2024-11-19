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
	// request.CreatedAt = time.Now()
	// request.UpdatedAt = time.Now()

	err := r.db.Create(&request).Error
	return err
}
