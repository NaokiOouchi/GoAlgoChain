package algo

// FindOperation locates items in a slice that match a predicate function.
// It returns all matching elements while preserving their original order.
type FindOperation[T any] struct {
	Predicate func(T) bool
}

// Apply performs the find operation on the data.
// It returns a new slice containing all elements that match the predicate.
//
// Example:
//
//	pipeline := NewPipeline[Product]().
//	    Find(func(p Product) bool { return p.Category == "Electronics" })
//	result, err := pipeline.Execute()
func (f *FindOperation[T]) Apply(data []T) ([]T, error) {
	var result []T
	for _, item := range data {
		if f.Predicate(item) {
			result = append(result, item)
		}
	}
	return result, nil
}

// Find adds a find operation to the pipeline.
// The predicate function should return true for items to be included in the result.
//
// Example:
//
//	pipeline.Find(func(order Order) bool {
//	    return order.Status == "Pending" && order.Total > 100
//	})
func (p *Pipeline[T]) Find(predicate func(T) bool) *Pipeline[T] {
	p.operations = append(p.operations, &FindOperation[T]{Predicate: predicate})
	return p
}
