package algo

// TakeOperation takes the first n items from the data.
type TakeOperation[T any] struct {
	Count int
}

// Apply applies the TakeOperation to a slice of data.
func (t *TakeOperation[T]) Apply(data []T) ([]T, error) {
	if t.Count <= 0 {
		return []T{}, nil
	}
	if t.Count > len(data) {
		t.Count = len(data)
	}
	takenData := data[:t.Count]
	return takenData, nil
}

// Take adds a TakeOperation to the pipeline.
func (p *Pipeline[T]) Take(count int) *Pipeline[T] {
	p.operations = append(p.operations, &TakeOperation[T]{Count: count})
	return p
}
