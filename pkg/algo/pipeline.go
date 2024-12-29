// Package algo provides a fluent API for chaining data operations and algorithms.
// It enables building data processing pipelines with a clean, chainable interface.
//
// Example usage:
//
//	pipeline := NewPipeline[int]().
//	    Filter(func(x int) bool { return x > 0 }).
//	    Map(func(x int) int { return x * 2 }).
//	    QuickSort(func(a, b int) bool { return a < b })
//
//	result, err := pipeline.Execute()
package algo

// Operation defines the interface for all pipeline operations.
// Each operation implements Apply to transform or process the data.
type Operation[T comparable] interface {
	Apply(data []T) ([]T, error)
}

// Pipeline represents a sequence of operations to be performed on data.
// Operations are executed in the order they were added to the pipeline.
type Pipeline[T comparable] struct {
	operations []Operation[T]
	data       []T
}

// NewPipeline creates a new Pipeline instance.
//
// Example:
//
//	pipeline := NewPipeline[User]()
func NewPipeline[T comparable]() *Pipeline[T] {
	return &Pipeline[T]{operations: []Operation[T]{}, data: []T{}}
}

// NewPipelineWithData creates a new Pipeline instance with initial data.
//
// Example:
//
//	data := []int{1, 2, 3, 4, 5}
//	pipeline := NewPipelineWithData(data).
//		Filter(func(x int) bool { return x > 2 }).
//		Map(func(x int) int { return x * 2 }).
//		QuickSort(func(a, b int) bool { return a < b })
//
//	result, err := pipeline.Execute()
func NewPipelineWithData[T comparable](data []T) *Pipeline[T] {
	return &Pipeline[T]{operations: []Operation[T]{}, data: data}
}

// WithData sets the initial data for the Pipeline.
// It can be used with NewPipeline to set data after pipeline creation.
//
// Example:
//
//	pipeline := NewPipeline[int]().
//		WithData([]int{1, 2, 3, 4, 5}).
//		Filter(func(x int) bool { return x > 2 }).
//		Map(func(x int) int { return x * 2 }).
//		QuickSort(func(a, b int) bool { return a < b })
//
//	result, err := pipeline.Execute()
func (p *Pipeline[T]) WithData(data []T) *Pipeline[T] {
	p.data = data
	return p
}

// AddOperation adds an operation to the pipeline.
// Returns the pipeline for method chaining.
func (p *Pipeline[T]) AddOperation(op Operation[T]) *Pipeline[T] {
	p.operations = append(p.operations, op)
	return p
}

// Execute runs all operations in the pipeline in sequence.
// Returns the final result or an error if any operation fails.
func (p *Pipeline[T]) Execute() ([]T, error) {
	var err error
	for _, op := range p.operations {
		p.data, err = op.Apply(p.data)
		if err != nil {
			return nil, err
		}
	}
	return p.data, nil
}

// GetOperations returns the slice of operations in the pipeline.
func (p *Pipeline[T]) GetOperations() []Operation[T] {
	return p.operations
}
