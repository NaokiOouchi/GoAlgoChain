package algo

import (
	"fmt"
)

// DistinctOperation removes duplicate items from the data.
type DistinctOperation[T any] struct {
	Equal func(a, b T) bool
}

// Apply applies the DistinctOperation to the data.
func (d *DistinctOperation[T]) Apply(data []T) ([]T, error) {
	distinctData := make([]T, 0)
	for _, item := range data {
		isDistinct := true
		for _, existing := range distinctData {
			if d.Equal(item, existing) {
				isDistinct = false
				break
			}
		}
		if isDistinct {
			distinctData = append(distinctData, item)
		}
	}
	fmt.Printf("DistinctOperation applied. %d unique items remaining.\n", len(distinctData))
	return distinctData, nil
}

// Distinct adds a DistinctOperation to the pipeline.
func (p *Pipeline[T]) Distinct(equal func(a, b T) bool) *Pipeline[T] {
	p.operations = append(p.operations, &DistinctOperation[T]{Equal: equal})
	return p
}
