package usecase

import (
	"errors"
	"time"

	"github.com/RE-PIT-BEM/re-pit-backend/internal/constant"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/model/domain"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/model/dto"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/services/request"
	"github.com/gin-gonic/gin"
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
		AcceptedBatch:               req.AcceptedBatch,

		WebsiteReleaseDate: req.WebsiteReleaseDate,
	}

	err := u.repo.Create(requestDomain)
	return err
}

func (u *RequestUsecase) GetAllRequest() ([]domain.Request, error) {
	requests, err := u.repo.FindAll()
	return requests, err
}

func (u *RequestUsecase) GetRequestByUserID(userID int) ([]domain.Request, error) {
	requests, err := u.repo.FindByUserID(userID)
	return requests, err
}

func (u *RequestUsecase) GetRequestByID(ctx *gin.Context, id string) (domain.Request, error) {
	userID, _ := ctx.Get("userId")
	role := ctx.GetString("role")

	request, err := u.repo.FindByID(id)

	if request.ID != userID && role != constant.ROLE_ADMIN {
		return domain.Request{}, errors.New("Is not your request!")
	}

	return request, err
}

func (u *RequestUsecase) AcceptRequest(id string) error {
	request, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}

	request.RequestStatus = constant.REQUEST_STATUS_ACCEPTED

	err = u.repo.Update(&request)
	return err
}

func (u *RequestUsecase) RejectRequest(id string, req dto.RejectRequestDTO) error {
	request, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}

	request.RequestStatus = constant.REQUEST_STATUS_REJECTED
	request.RejectedStatusMessage = req.RejectedStatusMessage

	err = u.repo.Update(&request)
	return err
}
