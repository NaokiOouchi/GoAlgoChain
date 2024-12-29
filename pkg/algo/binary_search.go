package algo

import (
	"fmt"
	"sort"
)

// BinarySearchOperation performs a binary search on the data.
// Predicate is a function that returns true if the item matches the target.
type BinarySearchOperation[T any] struct {
	Predicate  func(T) bool
	FoundIndex int
}

// Apply applies the binary search operation to the data.
// It returns the data and an error if the target is not found.
func (b *BinarySearchOperation[T]) Apply(data []T) ([]T, error) {
	index := sort.Search(len(data), func(i int) bool {
		return b.Predicate(data[i])
	})

	if index < len(data) && b.Predicate(data[index]) {
		b.FoundIndex = index
		return data, nil
	}

	return data, fmt.Errorf("target not found in data")
}

// BinarySearch adds a BinarySearchOperation to the pipeline.
func (p *Pipeline[T]) BinarySearch(predicate func(T) bool) *Pipeline[T] {
	p.operations = append(p.operations, &BinarySearchOperation[T]{Predicate: predicate})
	return p
}

// GetFoundIndex returns the index of the found item.
func (b *BinarySearchOperation[T]) GetFoundIndex() int {
	return b.FoundIndex
}
