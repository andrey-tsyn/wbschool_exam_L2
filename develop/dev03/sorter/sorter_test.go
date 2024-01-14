package sorting

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name           string
		input          []string
		exceptedResult []string
		exceptedError  error
		flags          map[string]interface{}
	}{
		{
			name:           "empty file",
			input:          []string{},
			exceptedResult: []string{},
			exceptedError:  nil,
			flags: map[string]interface{}{
				"reverse":    false,
				"numeric":    false,
				"unique":     false,
				"compColumn": 0,
			},
		},
		{
			name:           "default sorting",
			input:          []string{"42", "2", "34", "100", "string", "army"},
			exceptedResult: []string{"100", "2", "34", "42", "army", "string"},
			exceptedError:  nil,
			flags: map[string]interface{}{
				"reverse":    false,
				"numeric":    false,
				"unique":     false,
				"compColumn": 0,
			},
		},
		{
			name:           "reverse sorting",
			input:          []string{"B", "D", "C", "A "},
			exceptedResult: []string{"D", "C", "B", "A "},
			exceptedError:  nil,
			flags: map[string]interface{}{
				"reverse":    true,
				"numeric":    false,
				"unique":     false,
				"compColumn": 0,
			},
		},
		{
			name:           "numeric sorting",
			input:          []string{"42", "2", "34", "100"},
			exceptedResult: []string{"2", "34", "42", "100"},
			exceptedError:  nil,
			flags: map[string]interface{}{
				"reverse":    false,
				"numeric":    true,
				"unique":     false,
				"compColumn": 0,
			},
		},
		{
			name:           "numeric sorting with common strings",
			input:          []string{"42", "2", "34", "100", "string", "army"},
			exceptedResult: []string{"2", "34", "42", "100", "army", "string"},
			exceptedError:  nil,
			flags: map[string]interface{}{
				"reverse":    false,
				"numeric":    true,
				"unique":     false,
				"compColumn": 0,
			},
		},
		{
			name:           "unique sort",
			input:          []string{"42", "2", "34", "100", "100", "100", "string", "army", "army"},
			exceptedResult: []string{"100", "2", "34", "42", "army", "string"},
			exceptedError:  nil,
			flags: map[string]interface{}{
				"reverse":    false,
				"numeric":    false,
				"unique":     true,
				"compColumn": 0,
			},
		},
		{
			name:           "with comparing column",
			input:          []string{"42 a", "2 c", "34 b", "100 d", "army f", "string e"},
			exceptedResult: []string{"42 a", "34 b", "2 c", "100 d", "string e", "army f"},
			exceptedError:  nil,
			flags: map[string]interface{}{
				"reverse":    false,
				"numeric":    false,
				"unique":     false,
				"compColumn": 2,
			},
		},
		{
			name:           "numeric sorting with comparing column",
			input:          []string{"42 1", "2 3", "34 2", "100 4", "army 6", "string 5"},
			exceptedResult: []string{"42 1", "34 2", "2 3", "100 4", "string 5", "army 6"},
			exceptedError:  nil,
			flags: map[string]interface{}{
				"reverse":    false,
				"numeric":    false,
				"unique":     false,
				"compColumn": 2,
			},
		},
		{
			name:           "with comparing column and strings without needed column",
			input:          []string{"cat d", "filter", "ampule a", "dog", "medicine e"},
			exceptedResult: []string{"dog", "filter", "ampule a", "cat d", "medicine e"},
			exceptedError:  nil,
			flags: map[string]interface{}{
				"reverse":    false,
				"numeric":    false,
				"unique":     false,
				"compColumn": 2,
			},
		},
	}

	for _, test := range tests {
		sorter, err := NewStringSorter(
			WithReverse(test.flags["reverse"].(bool)),
			WithNumericSort(test.flags["numeric"].(bool)),
			WithUniqueOutput(test.flags["unique"].(bool)),
			WithComparingColumn(test.flags["compColumn"].(int)),
		)

		if err != nil {
			t.Errorf("sorter creation error occuried. Error: %s", err)
		}

		t.Run(test.name, func(t *testing.T) {
			result := sorter.Sort(test.input)

			if !reflect.DeepEqual(result, test.exceptedResult) {
				t.Errorf("unexpected result: %v, must be: %v", result, test.exceptedResult)
			}
		})
	}
}
