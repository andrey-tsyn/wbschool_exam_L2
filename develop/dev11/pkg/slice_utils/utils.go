package slice_utils

func Filter[T any](elems []T, predicate func(T) bool) []T {
	var result []T

	for _, elem := range elems {
		if predicate(elem) {
			result = append(result, elem)
		}
	}

	return result
}
