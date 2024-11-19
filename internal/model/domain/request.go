package domain

import "time"

type Request struct {
	ID     int `json:"id" gorm:"not null"`
	UserId int `json:"user_id" gorm:"not null"`

	ProgramName string `json:"program_name"`
	Department  string `json:"department"`
	Email       string `json:"email"`
	ContactName string `json:"contact_name"`
	GroupLink   string `json:"group_link"`

	ProgramDescription          string    `json:"program_description"`
	ProgramTimeline             string    `json:"program_timeline"`
	ProgramTimelineExtend       string    `json:"program_timeline_extend"`
	ProgramPhotoURL             string    `json:"program_photo_url"`
	ProgramLogoURL              string    `json:"program_logo_url"`
	ProgramDivision             string    `json:"program_division"`
	AcceptenceMessage           string    `json:"acceptence_message"`
	RejectionMessage            string    `json:"rejection_message"`
	ProgramQuotes               string    `json:"program_quotes"`
	ProgramRegistrationFlow     string    `json:"program_registration_flow"`
	ProgramApplicationForm      string    `json:"program_application_form"`
	ProgramRegistrationTemplate string    `json:"program_registration_template"`
	ProgramOpenDate             time.Time `json:"program_open_date"`
	ProgramCloseDate            time.Time `json:"program_close_date"`
	ProgramAnnouncementDate     time.Time `json:"program_announcement_date"`

	RequestStatus      string    `json:"status"`
	WebsiteReleaseDate time.Time `json:"website_release_date"`

	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}
