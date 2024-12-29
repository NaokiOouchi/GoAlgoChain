package algo

// QuickSortOperation sorts data using the quicksort algorithm.
// It provides efficient in-place sorting with O(n log n) average time complexity.
type QuickSortOperation[T any] struct {
	Comparator func(a, b T) bool
}

// Apply performs the quicksort operation on the data.
// It sorts the data in-place based on the provided comparator function.
//
// Example:
//
//	pipeline := NewPipeline[int]().
//	    QuickSort(func(a, b int) bool { return a < b })
//	result, err := pipeline.Execute()
func (q *QuickSortOperation[T]) Apply(data []T) ([]T, error) {
	if len(data) <= 1 {
		return data, nil
	}
	quickSortOptimized(data, 0, len(data)-1, q.Comparator)
	return data, nil
}

// quickSortOptimized is a helper function that sorts a slice of data using the QuickSort algorithm.
func quickSortOptimized[T any](data []T, low, high int, cmp func(a, b T) bool) {
	for low < high {
		if high-low <= 10 {
			insertionSort(data, low, high, cmp)
			break
		}
		pivot := medianOfThree(data, low, (low+high)/2, high, cmp)
		pi := partitionOptimized(data, low, high, pivot, cmp)

		if pi-low < high-pi {
			quickSortOptimized[T](data, low, pi-1, cmp)
			low = pi + 1
		} else {
			quickSortOptimized[T](data, pi+1, high, cmp)
			high = pi - 1
		}
	}
}

// insertionSort is a helper function that sorts a slice of data using the InsertionSort algorithm.
func medianOfThree[T any](data []T, low, mid, high int, cmp func(a, b T) bool) T {
	if cmp(data[mid], data[low]) {
		data[low], data[mid] = data[mid], data[low]
	}
	if cmp(data[high], data[low]) {
		data[low], data[high] = data[high], data[low]
	}
	if cmp(data[mid], data[high]) {
		data[mid], data[high] = data[high], data[mid]
	}
	return data[high]
}

// partitionOptimized is a helper function that partitions a slice of data using the QuickSort algorithm.
func partitionOptimized[T any](data []T, low, high int, pivot T, cmp func(a, b T) bool) int {
	left := low
	right := high - 1

	for {
		for left <= right && cmp(data[left], pivot) {
			left++
		}
		for left <= right && !cmp(data[right], pivot) {
			right--
		}
		if left > right {
			break
		}
		data[left], data[right] = data[right], data[left]
		left++
		right--
	}

	data[left], data[high] = data[high], data[left]
	return left
}

// insertionSort is a helper function that sorts a slice of data using the InsertionSort algorithm.
func insertionSort[T any](data []T, low, high int, cmp func(a, b T) bool) {
	for i := low + 1; i <= high; i++ {
		key := data[i]
		j := i - 1
		for j >= low && !cmp(data[j], key) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

// QuickSort adds a quicksort operation to the pipeline.
// The comparator function should return true when a should come before b in the sorted result.
//
// Example:
//
//	pipeline.QuickSort(func(a, b Product) bool {
//	    return a.Price < b.Price // Sort by price in ascending order
//	})
func (p *Pipeline[T]) QuickSort(comparator func(a, b T) bool) *Pipeline[T] {
	p.operations = append(p.operations, &QuickSortOperation[T]{Comparator: comparator})
	return p
}
