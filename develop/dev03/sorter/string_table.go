package sorting

import "strings"

type StringTable struct {
	table        [][]string
	comparingCol int
	numericSort  bool
}

func (s StringTable) Len() int {
	return len(s.table)
}

func (s StringTable) Less(i, j int) bool {
	currCol := s.comparingCol - 1

	iMissedColumn, jMissedColumn := len(s.table[i]) <= currCol, len(s.table[j]) <= currCol

	if iMissedColumn && jMissedColumn {
		return strings.Join(s.table[i], " ") < strings.Join(s.table[j], " ")
	} else if iMissedColumn {
		return true
	} else if jMissedColumn {
		return false
	}

	if s.numericSort {
		return compareStringsAsNumsIfPossible(s.table[i][currCol], s.table[j][currCol]) < 0
	}

	return s.table[i][currCol] < s.table[j][currCol]
}

func (s StringTable) Swap(i, j int) {
	s.table[i], s.table[j] = s.table[j], s.table[i]
}
