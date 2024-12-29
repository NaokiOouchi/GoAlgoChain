package algo

// MapOperation transforms each element in the data using a mapping function.
// The transformation is applied to all elements while preserving their order.
type MapOperation[T any] struct {
	Mapper func(T) T
}

// Apply performs the map operation on the data.
// It returns a new slice containing the transformed elements.
//
// Example:
//
//	pipeline := NewPipeline[int]().
//	    Map(func(x int) int { return x * 2 })
//	result, err := pipeline.Execute()
func (m *MapOperation[T]) Apply(data []T) ([]T, error) {
	mappedData := make([]T, len(data))
	for i := 0; i < len(data); i++ {
		mappedData[i] = m.Mapper(data[i])
	}
	return mappedData, nil
}

// Map adds a map operation to the pipeline.
// The mapper function defines how each element should be transformed.
//
// Example:
//
//	pipeline.Map(func(user User) User {
//	    user.Name = strings.ToUpper(user.Name)
//	    return user
//	})
func (p *Pipeline[T]) Map(mapper func(T) T) *Pipeline[T] {
	p.operations = append(p.operations, &MapOperation[T]{Mapper: mapper})
	return p
}
