package validation

import (
	"fmt"
)

type Error struct {
	msg string
}

func (e Error) Error() string {
	return e.msg
}

func NewValidationError(msg string) Error {
	return Error{msg}
}

func NewNegativeOrZeroIdError(fieldName string, providedValue int) Error {
	return Error{msg: fmt.Sprintf("%s can't be negative or zero, '%d' provided", fieldName, providedValue)}
}
