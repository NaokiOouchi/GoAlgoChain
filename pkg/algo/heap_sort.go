package algo

// HeapSortOperation sorts data using the heap sort algorithm.
// It provides stable sorting with O(n log n) time complexity.
type HeapSortOperation[T any] struct {
	Comparator func(a, b T) bool
}

// Apply performs the heap sort operation on the data.
// It sorts the data in-place based on the provided comparator function.
//
// Example:
//
//	pipeline := NewPipeline[int]().
//	    HeapSort(func(a, b int) bool { return a < b })
//	result, err := pipeline.Execute()
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
func buildMaxHeap[T any](data []T, cmp func(a, b T) bool) {
	n := len(data)
	for i := n/2 - 1; i >= 0; i-- {
		maxHeapify(data, i, n, cmp)
	}
}

// maxHeapify is a helper function that maintains the max heap property.
func maxHeapify[T any](data []T, i, n int, cmp func(a, b T) bool) {
	for {
		largest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < n && cmp(data[left], data[largest]) {
			largest = left
		}

		if right < n && cmp(data[right], data[largest]) {
			largest = right
		}

		if largest == i {
			break
		}

		data[i], data[largest] = data[largest], data[i]
		i = largest
	}
}

// HeapSort adds a heap sort operation to the pipeline.
// The comparator function should return true when a should come before b in the sorted result.
//
// Example:
//
//	pipeline.HeapSort(func(a, b User) bool {
//	    return a.Score > b.Score // Sort by score in descending order
//	})
func (p *Pipeline[T]) HeapSort(comparator func(a, b T) bool) *Pipeline[T] {
	p.operations = append(p.operations, &HeapSortOperation[T]{Comparator: comparator})
	return p
}
