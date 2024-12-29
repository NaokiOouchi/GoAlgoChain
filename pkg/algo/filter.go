package algo

// FilterOperation filters items in a slice of data based on a predicate.
type FilterOperation[T any] struct {
	Predicate func(T) bool
}

// Apply applies the FilterOperation to a slice of data.
func (f *FilterOperation[T]) Apply(data []T) ([]T, error) {
	filteredData := make([]T, 0, len(data))

	for i := 0; i < len(data); i++ {
		if f.Predicate(data[i]) {
			filteredData = append(filteredData, data[i])
		}
	}

	if cap(filteredData) > 2*len(data) {
		optimized := make([]T, len(filteredData))
		copy(optimized, filteredData)
		return optimized, nil
	}

	return filteredData, nil
}

// Filter adds a FilterOperation to the pipeline.
func (p *Pipeline[T]) Filter(predicate func(T) bool) *Pipeline[T] {
	p.operations = append(p.operations, &FilterOperation[T]{Predicate: predicate})
	return p
}
