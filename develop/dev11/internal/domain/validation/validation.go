package validation

import (
	"task11/internal/domain/model"
)

func ValidateEvent(event model.Event) error {
	if event.UserId < 1 {
		return NewNegativeOrZeroIdError("user_id", event.UserId)
	}
	if event.Date.IsZero() {
		return Error{
			msg: "date must be set",
		}
	}

	return nil
}
