package algo

// TakeOperation selects a specified number of elements from the beginning of the data.
// It preserves the order of the selected elements.
type TakeOperation[T any] struct {
	Count int
}

// Apply performs the take operation on the data.
// It returns a new slice containing the first Count elements.
//
// Example:
//
//	pipeline := NewPipeline[int]().
//	    Take(3) // Take first 3 elements
//	result, err := pipeline.Execute()
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

// Take adds a take operation to the pipeline.
// The count parameter specifies how many elements to select from the start.
//
// Example:
//
//	pipeline.Take(5) // Select first 5 elements
func (p *Pipeline[T]) Take(count int) *Pipeline[T] {
	p.operations = append(p.operations, &TakeOperation[T]{Count: count})
	return p
}
