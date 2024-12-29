package algo

import (
	"fmt"
)

// FilterOperation filters items in a slice of data based on a predicate.
type FilterOperation[T any] struct {
	Predicate func(T) bool
}

// Apply applies the FilterOperation to a slice of data.
func (f *FilterOperation[T]) Apply(data []T) ([]T, error) {
	filteredData := make([]T, 0)
	for _, item := range data {
		if f.Predicate(item) {
			filteredData = append(filteredData, item)
		}
	}
	fmt.Printf("FilterOperation applied. %d items remaining.\n", len(filteredData))
	return filteredData, nil
}

// Filter adds a FilterOperation to the pipeline.
func (p *Pipeline[T]) Filter(predicate func(T) bool) *Pipeline[T] {
	p.operations = append(p.operations, &FilterOperation[T]{Predicate: predicate})
	return p
}
