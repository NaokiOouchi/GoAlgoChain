package algo

import "fmt"

// LinearSearchOperation performs a linear search on the data.
type LinearSearchOperation[T any] struct {
	Predicate func(T) bool
}

// Apply applies the linear search operation to the data.
func (l *LinearSearchOperation[T]) Apply(data []T) ([]T, error) {
	for idx, item := range data {
		if l.Predicate(item) {
			fmt.Printf("Target found at index %d\n", idx)
			return data, nil
		}
	}
	return data, fmt.Errorf("target not found in data")
}

// LinearSearch adds a LinearSearchOperation to the pipeline.
func (p *Pipeline[T]) LinearSearch(predicate func(T) bool) *Pipeline[T] {
	p.operations = append(p.operations, &LinearSearchOperation[T]{Predicate: predicate})
	return p
}

// LinearSearchExact adds a LinearSearchOperation to the pipeline that searches for an exact match.
func (p *Pipeline[T]) LinearSearchExact(target T) *Pipeline[T] {
	p.operations = append(p.operations, &LinearSearchOperation[T]{Predicate: func(item T) bool {
		return item == target
	}})
	return p
}
