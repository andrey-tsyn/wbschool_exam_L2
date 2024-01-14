package sorting

import (
	"strconv"
	"strings"
	"unicode"
)

func compareStringsAsNumsIfPossible(a, b string) int {
	trimFunc := func(r rune) bool {
		return !unicode.IsDigit(r)
	}

	numsStrA := strings.TrimFunc(a, trimFunc)
	numsStrB := strings.TrimFunc(b, trimFunc)

	aNum, err := strconv.Atoi(numsStrA)
	if err != nil {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}

	bNum, err := strconv.Atoi(numsStrB)
	if err != nil {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}

	return aNum - bNum
}
