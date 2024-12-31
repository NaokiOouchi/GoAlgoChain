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
	if len(data) == 0 {
		return data, nil
	}
	distinctData := make([]T, 0, len(data))
	distinctData = append(distinctData, data[0])
	const batchSize = 64
	for i := 1; i < len(data); i++ {
		item := data[i]
		isDistinct := true
		for j := max(0, len(distinctData)-batchSize); j < len(distinctData); j++ {
			if d.Equal(item, distinctData[j]) {
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
