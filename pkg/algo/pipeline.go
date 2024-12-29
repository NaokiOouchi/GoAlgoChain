package algo

// Operation is an interface for applying an Operation to a slice of data.
type Operation[T comparable] interface {
	Apply(data []T) ([]T, error)
}

// Pipeline is a pipeline for applying multiple Operation to a slice of data.
type Pipeline[T comparable] struct {
	operations []Operation[T]
	data       []T
}

// NewPipeline creates a new Pipeline.
func NewPipeline[T comparable]() *Pipeline[T] {
	return &Pipeline[T]{operations: []Operation[T]{}, data: []T{}}
}

// AddOperation adds an Operation to the pipeline.
func (p *Pipeline[T]) AddOperation(op Operation[T]) *Pipeline[T] {
	p.operations = append(p.operations, op)
	return p
}

// Find adds a FindOperation to the pipeline.
func (p *Pipeline[T]) Find(predicate func(T) bool) *Pipeline[T] {
	p.operations = append(p.operations, &FindOperation[T]{Predicate: predicate})
	return p
}

// Execute executes the pipeline and returns the result.
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

func (p *Pipeline[T]) GetOperations() []Operation[T] {
	return p.operations
}
