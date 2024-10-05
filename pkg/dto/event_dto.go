package dto

import "time"

// EventCreateDTO - DTO для создания события
type EventCreateDTO struct {
	Title     string    `json:"title" binding:"required"`
	Details   string    `json:"details"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
}

// EventUpdateDTO - DTO для обновления события
type EventUpdateDTO struct {
	ID        uint      `json:"id" binding:"required"`
	Title     string    `json:"title" binding:"required"`
	Details   string    `json:"details"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
}

// EventResponseDTO - DTO для ответа с данными события
type EventResponseDTO struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Details   string    `json:"details"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
