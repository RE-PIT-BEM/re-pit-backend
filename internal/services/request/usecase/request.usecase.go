package usecase

import (
	"errors"
	"time"

	"github.com/RE-PIT-BEM/re-pit-backend/internal/model/domain"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/model/dto"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/services/request"
)

type RequestUsecase struct {
	repo request.RequestRepository
}

func NewRequestUsecase(repo request.RequestRepository) request.RequestUsecase {
	return &RequestUsecase{repo}
}

func (u *RequestUsecase) CreateRequest(req *dto.CreateRequestDTO, userId int) error {
	// check if website relasedate is less than 8 dyas from now
	if req.WebsiteReleaseDate.Before(time.Now().Add((time.Hour * (7 * 24)))) {
		return errors.New("Website Release Date must be at least 8 days from now")
	}

	//  check if open date is less than website release date
	if req.ProgramOpenDate.Before(req.WebsiteReleaseDate) {
		return errors.New("Program Open Date must be after Website Release Date")
	}

	// check if close date is less than open date
	if req.ProgramCloseDate.Before(req.ProgramOpenDate) {
		return errors.New("Program Close Date must be after Program Open Date")
	}

	// check if announcement date is less than close date
	if req.ProgramAnnouncementDate.Before(req.ProgramCloseDate) {
		return errors.New("Program Announcement Date must be after Program Close Date")
	}

	requestDomain := &domain.Request{
		UserId: userId,

		ProgramName: req.ProgramName,
		Department:  req.Department,
		ContactInfo: req.ContactInfo,
		ContactName: req.ContactName,
		GroupLink:   req.GroupLink,

		ProgramDescription:          req.ProgramDescription,
		ProgramTimeline:             req.ProgramTimeline,
		ProgramTimelineExtend:       req.ProgramTimelineExtend,
		ProgramPhotoURL:             req.ProgramPhotoURL,
		ProgramLogoURL:              req.ProgramLogoURL,
		ProgramDivision:             req.ProgramDivision,
		AcceptenceMessage:           req.AcceptenceMessage,
		RejectionMessage:            req.RejectionMessage,
		ProgramQuotes:               req.ProgramQuotes,
		ProgramRegistrationFlow:     req.ProgramRegistrationFlow,
		ProgramApplicationForm:      req.ProgramApplicationForm,
		ProgramRegistrationTemplate: req.ProgramRegistrationTemplate,
		ProgramOpenDate:             req.ProgramOpenDate,
		ProgramCloseDate:            req.ProgramCloseDate,
		ProgramAnnouncementDate:     req.ProgramAnnouncementDate,

		WebsiteReleaseDate: req.WebsiteReleaseDate,
	}

	err := u.repo.Create(requestDomain)
	return err
}
