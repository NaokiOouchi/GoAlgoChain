package algo

// QuickSortOperation is an Operation that sorts a slice of data using the QuickSort algorithm.
type QuickSortOperation[T any] struct {
	Comparator func(a, b T) int // Comparator(a, b) < 0: a < b; 0: a == b; >0: a > b
}

// Apply sorts the data using the QuickSort algorithm.
func (q *QuickSortOperation[T]) Apply(data []T) ([]T, error) {
	if len(data) == 0 {
		return data, nil
	}
	quickSort[T](data, 0, len(data)-1, q.Comparator)
	return data, nil
}

// quickSort sorts the data using the QuickSort algorithm.
func quickSort[T any](data []T, low, high int, cmp func(a, b T) int) {
	if low < high {
		pi := partition(data, low, high, cmp)
		quickSort[T](data, low, pi-1, cmp)
		quickSort[T](data, pi+1, high, cmp)
	}
}

// partition pivots the data around a pivot element and returns the index of the pivot.
func partition[T any](data []T, low, high int, cmp func(a, b T) int) int {
	pivot := data[high]
	i := low - 1
	for j := low; j < high; j++ {
		if cmp(data[j], pivot) < 0 {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}
	data[i+1], data[high] = data[high], data[i+1]
	return i + 1
}

// QuickSort adds a QuickSortOperation to the pipeline.
func (p *Pipeline[T]) QuickSort(comparator func(a, b T) int) *Pipeline[T] {
	p.operations = append(p.operations, &QuickSortOperation[T]{Comparator: comparator})
	return p
}
