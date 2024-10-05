package model

import "time"

// Event структура, представляющая событие
type Event struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `json:"title"`
	Details   string    `json:"details"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
