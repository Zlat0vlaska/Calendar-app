package service

import (
	"calendar-app/internal/repository"
	"calendar-app/pkg/model"
)

// EventService интерфейс для работы с событиями
type EventService interface {
	CreateEvent(event *model.Event) error
	UpdateEvent(event *model.Event) error
	DeleteEvent(id uint) error
	GetEventByID(id uint) (*model.Event, error)
	GetAllEvents() ([]model.Event, error)
}

// eventServiceImpl реализация интерфейса EventService
type eventServiceImpl struct {
	eventRepo repository.EventRepository
}

func NewEventService(eventRepo repository.EventRepository) EventService {
	return &eventServiceImpl{eventRepo: eventRepo}
}

func (s *eventServiceImpl) CreateEvent(event *model.Event) error {
	return s.eventRepo.Create(event)
}

func (s *eventServiceImpl) UpdateEvent(event *model.Event) error {
	return s.eventRepo.Update(event)
}

func (s *eventServiceImpl) DeleteEvent(id uint) error {
	return s.eventRepo.Delete(id)
}

func (s *eventServiceImpl) GetEventByID(id uint) (*model.Event, error) {
	return s.eventRepo.FindByID(id)
}

func (s *eventServiceImpl) GetAllEvents() ([]model.Event, error) {
	return s.eventRepo.FindAll()
}
