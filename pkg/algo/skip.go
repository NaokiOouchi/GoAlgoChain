package algo

// SkipOperation skips the first n items from the data.
type SkipOperation[T any] struct {
	Count int
}

// Apply applies the SkipOperation to a slice of data.
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

// Skip adds a SkipOperation to the pipeline.
func (p *Pipeline[T]) Skip(count int) *Pipeline[T] {
	p.operations = append(p.operations, &SkipOperation[T]{Count: count})
	return p
}
