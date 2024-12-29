package algo

// HeapSortOperation is an Operation that sorts a slice of data using the HeapSort algorithm.
type HeapSortOperation[T any] struct {
	Comparator func(a, b T) int // Comparator(a, b) < 0: a < b; 0: a == b; >0: a > b
}

// Apply is an Operation that sorts a slice of data using the HeapSort algorithm.
func (h *HeapSortOperation[T]) Apply(data []T) ([]T, error) {
	if len(data) == 0 {
		return data, nil
	}
	buildMaxHeap(data, h.Comparator)
	for i := len(data) - 1; i > 0; i-- {
		data[0], data[i] = data[i], data[0]
		maxHeapify(data, 0, i, h.Comparator)
	}
	return data, nil
}

// buildMaxHeap is a helper function that builds a max heap from a slice of data.
func buildMaxHeap[T any](data []T, cmp func(a, b T) int) {
	n := len(data)
	for i := n/2 - 1; i >= 0; i-- {
		maxHeapify(data, i, n, cmp)
	}
}

// maxHeapify is a helper function that maintains the max heap property.
func maxHeapify[T any](data []T, i, n int, cmp func(a, b T) int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && cmp(data[left], data[largest]) > 0 {
		largest = left
	}

	if right < n && cmp(data[right], data[largest]) > 0 {
		largest = right
	}

	if largest != i {
		data[i], data[largest] = data[largest], data[i]
		maxHeapify[T](data, largest, n, cmp)
	}
}

// HeapSort adds a HeapSortOperation to the pipeline.
func (p *Pipeline[T]) HeapSort(comparator func(a, b T) int) *Pipeline[T] {
	p.operations = append(p.operations, &HeapSortOperation[T]{Comparator: comparator})
	return p
}
