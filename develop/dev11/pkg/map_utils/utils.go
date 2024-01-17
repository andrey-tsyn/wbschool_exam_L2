package map_utils

func GetKeys[K comparable, V any](m map[K]V) []V {
	result := make([]V, len(m))
	for _, v := range m {
		result = append(result, v)
	}

	return result
}
