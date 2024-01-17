package service

import (
	"task11/internal/domain/model"
	"task11/internal/domain/repository"
	"task11/internal/domain/validation"
	"time"
)

type EventService struct {
	repository repository.EventRepository
}

func NewEventService(repository repository.EventRepository) *EventService {
	return &EventService{repository: repository}
}

func (s *EventService) Update(event model.Event) error {
	if event.Id < 1 {
		return validation.NewNegativeOrZeroIdError("id", event.Id)
	}
	err := validation.ValidateEvent(event)
	if err != nil {
		return err
	}

	return s.repository.Update(event)
}

func (s *EventService) Create(event model.Event) (int, error) {
	err := validation.ValidateEvent(event)
	if err != nil {
		return -1, err
	}

	return s.repository.Create(event)
}

func (s *EventService) Delete(id int) error {
	if id < 1 {
		return validation.NewNegativeOrZeroIdError("id", id)
	}
	return s.repository.Delete(id)
}

func (s *EventService) GetByUserId(userId int, from, to time.Time) ([]model.Event, error) {
	return s.repository.GetByUserId(userId, from, to)
}
