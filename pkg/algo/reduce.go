package algo

import (
	"fmt"
)

// ReduceOperation aggregates all elements in the data into a single value.
// It applies the reducer function sequentially from left to right.
type ReduceOperation[T any] struct {
	Reducer func(acc, item T) T
}

// Apply performs the reduce operation on the data.
// It returns a slice containing the single accumulated result.
// Returns an error if the input slice is empty.
//
// Example:
//
//	pipeline := NewPipeline[int]().
//	    Reduce(func(acc, item int) int { return acc + item })
//	result, err := pipeline.Execute() // Sums all numbers
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

// Reduce adds a reduce operation to the pipeline.
// The reducer function combines the accumulator with each item to produce a new accumulator.
//
// Example:
//
//	pipeline.Reduce(func(acc, item Order) Order {
//	    acc.TotalAmount += item.Amount
//	    return acc
//	})
func (p *Pipeline[T]) Reduce(reducer func(acc, item T) T) *Pipeline[T] {
	p.operations = append(p.operations, &ReduceOperation[T]{Reducer: reducer})
	return p
}
