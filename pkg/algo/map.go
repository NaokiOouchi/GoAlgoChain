package algo

// MapOperation applies a mapping function to each item in the data.
type MapOperation[T any] struct {
	Mapper func(T) T
}

// Apply applies the MapOperation to the data.
func (m *MapOperation[T]) Apply(data []T) ([]T, error) {
	mappedData := make([]T, len(data))
	for i, item := range data {
		mappedData[i] = m.Mapper(item)
	}
	return mappedData, nil
}

// Map adds a MapOperation to the pipeline.
func (p *Pipeline[T]) Map(mapper func(T) T) *Pipeline[T] {
	p.operations = append(p.operations, &MapOperation[T]{Mapper: mapper})
	return p
}
