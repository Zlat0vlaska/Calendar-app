package dto

import "time"

// EventDTO - DTO
type EventDTO struct {
	ID              uint      `json:"primaryKey"`
	Name            string    `json:"name"`
	Place           string    `json:"place"`
	Comment         *string   `json:"comment"`
	RecipientEmails []*string `json:"recipient_emails"`
	ApplicantEmail  string    `json:"applicant_email"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	IsFullDay       bool      `json:"is_full_day"`
	IsOnline        bool      `json:"is_online"`
	AuthorEmail     string    `json:"author_email"`
}

// EventResponseDTO - DTO для ответа с данными события
type EventResponseDTO struct {
	ID              uint      `json:"primaryKey"`
	Name            string    `json:"name"`
	Place           string    `json:"place"`
	Comment         *string   `json:"comment"`
	RecipientEmails []*string `json:"recipient_emails"`
	ApplicantEmail  string    `json:"applicant_email"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	IsFullDay       bool      `json:"is_full_day"`
	IsOnline        bool      `json:"is_online"`
	AuthorEmail     string    `json:"author_email"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
