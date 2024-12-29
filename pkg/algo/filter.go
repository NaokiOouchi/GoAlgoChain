package algo

// FilterOperation filters items in a slice based on a predicate function.
// It preserves the order of items that match the predicate.
type FilterOperation[T any] struct {
	Predicate func(T) bool
}

// Apply performs the filter operation on the data.
// It returns a new slice containing only the elements that satisfy the predicate.
//
// Example:
//
//	pipeline := NewPipeline[User]().
//	    Filter(func(u User) bool { return u.Age >= 18 })
//	result, err := pipeline.Execute()
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

// Filter adds a filter operation to the pipeline.
// The predicate function should return true for items to keep in the result.
//
// Example:
//
//	pipeline.Filter(func(item Product) bool {
//	    return item.Price > 1000 && item.InStock
//	})
func (p *Pipeline[T]) Filter(predicate func(T) bool) *Pipeline[T] {
	p.operations = append(p.operations, &FilterOperation[T]{Predicate: predicate})
	return p
}
