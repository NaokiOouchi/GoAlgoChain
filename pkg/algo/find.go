package algo

// FindOperation finds items in a slice of data that match a predicate.
type FindOperation[T any] struct {
	Predicate func(T) bool
}

// Apply applies the FindOperation to a slice of data.
func (f *FindOperation[T]) Apply(data []T) ([]T, error) {
	var result []T
	for _, item := range data {
		if f.Predicate(item) {
			result = append(result, item)
		}
	}
	return result, nil
}
