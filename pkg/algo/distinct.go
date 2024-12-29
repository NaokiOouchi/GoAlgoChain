package algo

// DistinctOperation removes duplicate items from the data based on the provided equality function.
// It preserves the order of first occurrence of each unique item.
type DistinctOperation[T any] struct {
	Equal func(a, b T) bool
}

// Apply performs the distinct operation on the data.
// It returns a new slice containing only unique elements based on the Equal function.
//
// Example:
//
//	pipeline := NewPipeline[Item]().
//	    Distinct(func(a, b Item) bool { return a.ID == b.ID })
//	result, err := pipeline.Execute()
func (d *DistinctOperation[T]) Apply(data []T) ([]T, error) {
	distinctData := make([]T, 0)
	for _, item := range data {
		isDistinct := true
		for _, existing := range distinctData {
			if d.Equal(item, existing) {
				isDistinct = false
				break
			}
		}
		if isDistinct {
			distinctData = append(distinctData, item)
		}
	}
	return distinctData, nil
}

// Distinct adds a distinct operation to the pipeline.
// The equal function should return true when two items are considered equal.
//
// Example:
//
//	pipeline.Distinct(func(a, b User) bool {
//	    return a.ID == b.ID && a.Role == b.Role
//	})
func (p *Pipeline[T]) Distinct(equal func(a, b T) bool) *Pipeline[T] {
	p.operations = append(p.operations, &DistinctOperation[T]{Equal: equal})
	return p
}
