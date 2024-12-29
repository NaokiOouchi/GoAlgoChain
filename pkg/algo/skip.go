package algo

// SkipOperation bypasses a specified number of elements from the beginning of the data.
// It returns the remaining elements while preserving their order.
type SkipOperation[T any] struct {
	Count int
}

// Apply performs the skip operation on the data.
// It returns a new slice excluding the first Count elements.
//
// Example:
//
//	pipeline := NewPipeline[string]().
//	    Skip(2) // Skip first 2 elements
//	result, err := pipeline.Execute()
func (s *SkipOperation[T]) Apply(data []T) ([]T, error) {
	if s.Count <= 0 {
		return data, nil
	}
	if s.Count >= len(data) {
		return []T{}, nil
	}
	skippedData := data[s.Count:]
	return skippedData, nil
}

// Skip adds a skip operation to the pipeline.
// The count parameter specifies how many elements to skip from the start.
//
// Example:
//
//	pipeline.Skip(5) // Skip first 5 elements
func (p *Pipeline[T]) Skip(count int) *Pipeline[T] {
	p.operations = append(p.operations, &SkipOperation[T]{Count: count})
	return p
}
