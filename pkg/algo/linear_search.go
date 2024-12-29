package algo

import "fmt"

// LinearSearchOperation performs a sequential search through the data.
// It searches for elements that match the given predicate function.
type LinearSearchOperation[T any] struct {
	Predicate func(T) bool
}

// Apply performs the linear search operation on the data.
// It returns the data and an error if no matching element is found.
//
// Example:
//
//	pipeline := NewPipeline[Product]().
//	    LinearSearch(func(p Product) bool { return p.SKU == "ABC123" })
//	result, err := pipeline.Execute()
func (l *LinearSearchOperation[T]) Apply(data []T) ([]T, error) {
	for _, item := range data {
		if l.Predicate(item) {
			return data, nil
		}
	}
	return data, fmt.Errorf("target not found in data")
}

// LinearSearch adds a linear search operation to the pipeline.
// The predicate function should return true when the target element is found.
//
// Example:
//
//	pipeline.LinearSearch(func(user User) bool {
//	    return user.Email == "example@email.com"
//	})
func (p *Pipeline[T]) LinearSearch(predicate func(T) bool) *Pipeline[T] {
	p.operations = append(p.operations, &LinearSearchOperation[T]{Predicate: predicate})
	return p
}

// LinearSearchExact adds a linear search operation that looks for an exact match.
//
// Example:
//
//	pipeline.LinearSearchExact(targetUser)
func (p *Pipeline[T]) LinearSearchExact(target T) *Pipeline[T] {
	p.operations = append(p.operations, &LinearSearchOperation[T]{Predicate: func(item T) bool {
		return item == target
	}})
	return p
}
