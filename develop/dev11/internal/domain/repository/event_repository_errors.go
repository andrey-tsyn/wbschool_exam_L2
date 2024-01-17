package repository

import "fmt"

type EventNotFoundError struct {
	EventId int
}

func (e EventNotFoundError) Error() string {
	return fmt.Sprintf("event with id '%d' not found", e.EventId)
}

type EventAlreadyExist struct {
	EventId int
}

func (e EventAlreadyExist) Error() string {
	return fmt.Sprintf("event with id '%d' already exist", e.EventId)
}
