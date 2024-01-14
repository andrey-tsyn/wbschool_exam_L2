package slice_utils

// RemoveDuplicates удаляет одинаковые элементы из слайса.
func RemoveDuplicates[T comparable](slice []T) []T {
	m := make(map[T]struct{})

	resultList := make([]T, 0, len(slice))

	for _, val := range slice {
		if _, ok := m[val]; !ok {
			resultList = append(resultList, val)
			m[val] = struct{}{}
		}
	}
	return resultList
}
