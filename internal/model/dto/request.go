package dto

import "time"

type CreateRequestDTO struct {
	ProgramName string `json:"program_name" binding:"required"`
	Department  string `json:"department" binding:"required"`
	ContactInfo string `json:"contact_info" binding:"required"`
	ContactName string `json:"contact_name" binding:"required"`
	GroupLink   string `json:"group_link" binding:"required"`

	ProgramDescription          string    `json:"program_description" binding:"required"`
	ProgramTimeline             string    `json:"program_timeline" binding:"required"`
	ProgramTimelineExtend       string    `json:"program_timeline_extend"`
	ProgramPhotoURL             string    `json:"program_photo_url" binding:"required,url"`
	ProgramLogoURL              string    `json:"program_logo_url"`
	ProgramDivision             string    `json:"program_division" binding:"required"`
	AcceptenceMessage           string    `json:"acceptence_message" binding:"required"`
	RejectionMessage            string    `json:"rejection_message" binding:"required"`
	ProgramQuotes               string    `json:"program_quotes" binding:"required"`
	ProgramRegistrationFlow     string    `json:"program_registration_flow" binding:"required"`
	ProgramApplicationForm      string    `json:"program_application_form" binding:"required"`
	ProgramRegistrationTemplate string    `json:"program_registration_template" binding:"required"`
	ProgramOpenDate             time.Time `json:"program_open_date" binding:"required"`
	ProgramCloseDate            time.Time `json:"program_close_date" binding:"required"`
	ProgramAnnouncementDate     time.Time `json:"program_announcement_date" binding:"required"`

	WebsiteReleaseDate time.Time `json:"website_release_date" binding:"required"`
}

type RejectRequestDTO struct {
	RejectedStatusMessage string `json:"rejected_status_message" binding:"required"`
}
