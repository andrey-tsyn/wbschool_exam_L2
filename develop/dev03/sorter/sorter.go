package sorting

import (
	"errors"
	"fmt"
	"slices"
	"sort"
	"strings"
	"task3/slice_utils"
)

// StringSorter sorts string slice with many options.
//
// Fields:
//
// comparingColumn - number of column to compare, default separator is " ", numeration starts from 1
type StringSorter struct {
	reverse         bool
	comparingColumn int
	uniqueOutput    bool
	numericSort     bool
}

// Код, связанный с конструктором

type option func(s *StringSorter) error

func NewStringSorter(options ...option) (*StringSorter, error) {
	sorter := &StringSorter{
		reverse:         false,
		comparingColumn: 0,
		uniqueOutput:    false,
		numericSort:     false,
	}

	for _, opt := range options {
		if err := opt(sorter); err != nil {
			return nil, err
		}
	}

	return sorter, nil
}

func WithReverse(r bool) option {
	return func(s *StringSorter) error {
		s.reverse = r
		return nil
	}
}

func WithUniqueOutput(u bool) option {
	return func(s *StringSorter) error {
		s.uniqueOutput = u
		return nil
	}
}

func WithComparingColumn(col int) option {
	return func(s *StringSorter) error {
		if col < 0 {
			return errors.New(fmt.Sprintf("column number can't be negative, %d provided", col))
		}
		s.comparingColumn = col
		return nil
	}
}

func WithNumericSort(n bool) option {
	return func(s *StringSorter) error {
		s.numericSort = n
		return nil
	}
}

// Код с логикой

func (s *StringSorter) Sort(strs []string) []string {
	sortMethod := s.getSortingMethod()

	strs = sortMethod(strs)

	if s.uniqueOutput {
		strs = slice_utils.RemoveDuplicates(strs)
	}

	if s.reverse {
		slices.Reverse(strs)
	}

	return strs
}

func (s *StringSorter) getSortingMethod() func([]string) []string {
	if s.comparingColumn == 0 {
		return s.basicSort
	} else {
		return s.columnSort
	}
}

// basicSort использует стандартную сортировку строк из пакета sort
func (s *StringSorter) basicSort(strs []string) []string {
	if s.numericSort {
		slices.SortFunc(strs, compareStringsAsNumsIfPossible)
		return strs
	}

	sort.Strings(strs)
	return strs
}

func (s *StringSorter) columnSort(strs []string) []string {
	stringTable := StringTable{
		table:        make([][]string, len(strs)),
		comparingCol: s.comparingColumn,
		numericSort:  s.numericSort,
	}

	for i, val := range strs {
		stringTable.table[i] = strings.Split(val, " ")
	}

	sort.Sort(stringTable)

	for i, val := range stringTable.table {
		strs[i] = strings.Join(val, " ")
	}

	return strs
}
