package slices

// Filter returns a sub-slice of items, where all the items for which the
// filter function returned false, have been removed.
// The original sorting is preserved.
func Filter[T any](items []T, filter func(i T) bool) []T {
	filtered := make([]T, 0)
	for _, i := range items {
		if filter(i) {
			filtered = append(filtered, i)
		}
	}
	return filtered
}
