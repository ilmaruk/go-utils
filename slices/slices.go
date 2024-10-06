package slices

func Filter[T any](items []T, filter func(i T) bool) []T {
	filtered := make([]T, 0)
	for _, i := range items {
		if filter(i) {
			filtered = append(filtered, i)
		}
	}
	return filtered
}
