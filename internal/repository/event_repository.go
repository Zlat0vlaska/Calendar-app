package repository

import (
	"calendar-app/pkg/model"

	"gorm.io/gorm"
)

// EventRepository интерфейс для работы с событиями
type EventRepository interface {
	Create(event *model.Event) error
	Update(event *model.Event) error
	Delete(id uint) error
	FindByID(id uint) (*model.Event, error)
	FindAll() ([]model.Event, error)
}

// eventRepositoryImpl реализация интерфейса EventRepository
type eventRepositoryImpl struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepositoryImpl{db: db}
}

func (r *eventRepositoryImpl) Create(event *model.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepositoryImpl) Update(event *model.Event) error {
	return r.db.Save(event).Error
}

func (r *eventRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&model.Event{}, id).Error
}

func (r *eventRepositoryImpl) FindByID(id uint) (*model.Event, error) {
	var event model.Event
	if err := r.db.First(&event, id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepositoryImpl) FindAll() ([]model.Event, error) {
	var events []model.Event
	if err := r.db.Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}
