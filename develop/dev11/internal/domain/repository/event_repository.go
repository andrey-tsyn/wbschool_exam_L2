package repository

import (
	"task11/internal/domain/model"
	"time"
)

type EventRepository interface {
	// Update updates event, if exist return error.
	Update(event model.Event) error
	// Create creates event and returns event id, if exist returns error.
	Create(event model.Event) (int, error)
	// Delete deletes event. If doesn't exist returns error.
	Delete(id int) error
	// GetByUserId returns events for time interval
	GetByUserId(userId int, from, to time.Time) ([]model.Event, error)
}
