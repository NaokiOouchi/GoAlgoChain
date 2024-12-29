package algo

import (
	"fmt"
	"sort"
)

// BinarySearchOperation performs a binary search on sorted data.
// It requires the data to be sorted according to the predicate function's ordering.
type BinarySearchOperation[T any] struct {
	Predicate  func(T) bool
	FoundIndex int
}

// Apply performs the binary search operation on the data.
// It returns the original data slice and an error if the target is not found.
// The operation expects the data to be pre-sorted for correct results.
//
// Example:
//
//	pipeline := NewPipeline[int]().
//	    QuickSort(func(a, b int) bool { return a < b }).
//	    BinarySearch(func(a int) bool { return a == 42 })
//	result, err := pipeline.Execute()
func (b *BinarySearchOperation[T]) Apply(data []T) ([]T, error) {
	b.FoundIndex = -1
	index := sort.Search(len(data), func(i int) bool {
		return b.Predicate(data[i])
	})

	if index < len(data) && b.Predicate(data[index]) {
		b.FoundIndex = index
		return data, nil
	}

	return data, fmt.Errorf("target not found in data")
}

// BinarySearch adds a binary search operation to the pipeline.
// The predicate function should return true when the target element is found.
//
// Example:
//
//	pipeline.BinarySearch(func(item int) bool { return item == targetValue })
func (p *Pipeline[T]) BinarySearch(predicate func(T) bool) *Pipeline[T] {
	p.operations = append(p.operations, &BinarySearchOperation[T]{Predicate: predicate})
	return p
}

// GetFoundIndex returns the index of the found element after a successful binary search.
// Returns -1 if the element was not found or if the search has not been executed.
func (b *BinarySearchOperation[T]) GetFoundIndex() int {
	return b.FoundIndex
}
