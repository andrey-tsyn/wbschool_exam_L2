package main

import (
	"reflect"
	"testing"
)

func TestGetAnagrams(t *testing.T) {
	tests := []struct {
		name           string
		input          []string
		expectedResult map[string][]string
		expectedError  error
	}{
		{
			name:  "with common lower strings",
			input: []string{"пятка", "пятак", "тяпка", "листок", "слиток", "столик"},
			expectedResult: map[string][]string{
				"акптя":  {"пятак", "пятка", "тяпка"},
				"иклост": {"листок", "слиток", "столик"},
			},
			expectedError: nil,
		},
		{
			name:  "with common strings",
			input: []string{"пЯтка", "пятак", "тяПка", "листоК", "СЛИток", "столик"},
			expectedResult: map[string][]string{
				"акптя":  {"пятак", "пятка", "тяпка"},
				"иклост": {"листок", "слиток", "столик"},
			},
			expectedError: nil,
		},
		{
			name:           "empty",
			input:          []string{},
			expectedResult: map[string][]string{},
			expectedError:  nil,
		},
		{
			name:  "with common strings and words without anagrams",
			input: []string{"пЯтка", "пятак", "тяПка", "листоК", "СЛИток", "столик", "Яна", "Министр"},
			expectedResult: map[string][]string{
				"акптя":  {"пятак", "пятка", "тяпка"},
				"иклост": {"листок", "слиток", "столик"},
			},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetAnagrams(test.input)

			if !reflect.DeepEqual(result, test.expectedResult) {
				t.Errorf("")
			}
		})
	}
}
