package repository

import (
	"calendar-app/pkg/model"
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
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
	db *sql.DB
}

func NewEventRepository(db *sql.DB) EventRepository {
	return &eventRepositoryImpl{db: db}
}

func (r *eventRepositoryImpl) Create(event *model.Event) error {
	sql, args, err := sq.Insert("events").
		Columns("name", "place", "comment", "recipientEmails", "applicantEmail", "startDate", "endDate", "isFullDay", "isOnline", "authorEmail").
		Values(event.Name, event.Place, event.Comment, event.RecipientEmails, event.ApplicantEmail, event.StartDate, event.EndDate, event.IsFullDay, event.IsOnline, event.AuthorEmail).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(context.Background(), sql, args...)
	return err
}

func (r *eventRepositoryImpl) Update(event *model.Event) error {
	sql, args, err := sq.Update("events").
		Set("name", event.Name).
		Set("place", event.Place).
		Set("comment", event.Comment).
		Set("recipientEmails", event.RecipientEmails).
		Set("applicantEmail", event.ApplicantEmail).
		Set("startDate", event.StartDate).
		Set("endDate", event.EndDate).
		Set("IsFullDay", event.IsFullDay).
		Set("IsOnline", event.IsOnline).
		Set("authorEmail", event.AuthorEmail).
		Where(sq.Eq{"id": event.ID}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(context.Background(), sql, args...)
	return err
}

func (r *eventRepositoryImpl) Delete(id uint) error {
	sql, args, err := sq.Delete("events").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(context.Background(), sql, args...)
	return err
}

func (r *eventRepositoryImpl) FindByID(id uint) (*model.Event, error) {
	var event model.Event

	sql, args, err := sq.Select("*").
		From("events").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	row := r.db.QueryRowContext(context.Background(), sql, args...)

	if err := row.Scan(&event.ID, &event.Name, &event.Place, &event.Comment,
		&event.RecipientEmails, &event.ApplicantEmail, &event.StartDate, &event.EndDate, &event.IsFullDay, &event.IsOnline, &event.AuthorEmail); err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *eventRepositoryImpl) FindAll() ([]model.Event, error) {
	var events []model.Event

	sql, args, err := sq.Select("*").
		From("events").
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(context.Background(), sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var event model.Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Place, &event.Comment,
			&event.RecipientEmails, &event.ApplicantEmail, &event.StartDate, &event.EndDate, &event.IsFullDay, &event.IsOnline, &event.AuthorEmail); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
