package domain

import "time"

type Request struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserId int  `json:"user_id" gorm:"not null"`

	ProgramName string `json:"program_name" gorm:"not null"`
	Department  string `json:"department" gorm:"not null;type:varchar(100)"`
	ContactName string `json:"contact_name" gorm:"not null;type:varchar(100)"`
	ContactInfo string `json:"contact_info" gorm:"not null;type:varchar(100)"`
	GroupLink   string `json:"group_link" gorm:"not null;type:varchar(100)"`

	ProgramDescription          string    `json:"program_description" gorm:"not null"`
	ProgramTimeline             string    `json:"program_timeline" gorm:"not null"`
	ProgramTimelineExtend       string    `json:"program_timeline_extend" gorm:"not null"`
	ProgramPhotoURL             string    `json:"program_photo_url" gorm:"not null;type:varchar(100)"`
	ProgramLogoURL              string    `json:"program_logo_url" gorm:"not null;type:varchar(100)"`
	ProgramDivision             string    `json:"program_division" gorm:"not null;type:varchar(100)"`
	AcceptenceMessage           string    `json:"acceptence_message" gorm:"not null"`
	RejectionMessage            string    `json:"rejection_message" gorm:"not null"`
	ProgramQuotes               string    `json:"program_quotes" gorm:"not null"`
	ProgramRegistrationFlow     string    `json:"program_registration_flow" gorm:"not null"`
	ProgramApplicationForm      string    `json:"program_application_form" gorm:"not null"`
	ProgramRegistrationTemplate string    `json:"program_registration_template" gorm:"not null"`
	ProgramOpenDate             time.Time `json:"program_open_date" gorm:"not null"`
	ProgramCloseDate            time.Time `json:"program_close_date" gorm:"not null"`
	ProgramAnnouncementDate     time.Time `json:"program_announcement_date" gorm:"not null"`

	RequestStatus      string    `json:"status" gorm:"not null"`
	WebsiteReleaseDate time.Time `json:"website_release_date" gorm:"not null"`

	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}
