package model

import "time"

// Event структура, представляющая событие
type Event struct {
	ID              uint      `db:"primaryKey"`
	Name            string    `db:"name"`
	Place           string    `db:"place"`
	Comment         *string   `db:"comment"`
	RecipientEmails []*string `db:"recipient_emails"`
	ApplicantEmail  string    `db:"applicant_email"`
	StartDate       time.Time `db:"start_date"`
	EndDate         time.Time `db:"end_date"`
	IsFullDay       bool      `db:"is_full_day"`
	IsOnline        bool      `db:"is_online"`
	AuthorEmail     string    `db:"author_email"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
