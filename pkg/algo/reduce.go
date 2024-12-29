package algo

import (
	"fmt"
)

// ReduceOperation reduces a slice of data to a single value.
type ReduceOperation[T any] struct {
	Reducer func(acc, item T) T
}

// Apply applies the ReduceOperation to the data.
func (r *ReduceOperation[T]) Apply(data []T) ([]T, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("ReduceOperation: cannot reduce an empty slice")
	}

	acc := data[0]
	for _, item := range data[1:] {
		acc = r.Reducer(acc, item)
	}
	return []T{acc}, nil
}

// Reduce adds a ReduceOperation to the pipeline.
func (p *Pipeline[T]) Reduce(reducer func(acc, item T) T) *Pipeline[T] {
	p.operations = append(p.operations, &ReduceOperation[T]{Reducer: reducer})
	return p
}
