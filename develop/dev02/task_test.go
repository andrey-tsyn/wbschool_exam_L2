package main

import (
	"errors"
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		exceptedResult string
		exceptedError  error
	}{
		{
			name:           "Empty string",
			input:          "",
			exceptedResult: "",
			exceptedError:  nil,
		},
		{
			name:           "One character",
			input:          "c",
			exceptedResult: "c",
			exceptedError:  nil,
		},
		{
			name:           "Digits after char",
			input:          "c4",
			exceptedResult: "cccc",
			exceptedError:  nil,
		},
		{
			name:           "Digits after char several times",
			input:          "a1b2c3d4e5fff",
			exceptedResult: "abbcccddddeeeeefff",
			exceptedError:  nil,
		},
		{
			name:           "With escape sequence",
			input:          "qwe\\4\\5",
			exceptedResult: "qwe45",
			exceptedError:  nil,
		},
		{
			name:           "With escape sequence 2",
			input:          "qwe\\45",
			exceptedResult: "qwe44444",
			exceptedError:  nil,
		},
		{
			name:           "With escape sequence 3",
			input:          "qwe\\\\5",
			exceptedResult: "qwe\\\\\\\\\\",
			exceptedError:  nil,
		},
		{
			name:           "Starts with digit",
			input:          "32k",
			exceptedResult: "",
			exceptedError:  IncorrectString,
		},
		{
			name:           "Digits only",
			input:          "32461237",
			exceptedResult: "",
			exceptedError:  IncorrectString,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := UnpackString(test.input)
			if err != nil {
				if errors.Is(err, test.exceptedError) {
					return
				}
				t.Errorf("unexpected error recieved")
				return
			}

			if result != test.exceptedResult {
				t.Errorf("unexpected result recieved.")
			}
		})
	}
}
