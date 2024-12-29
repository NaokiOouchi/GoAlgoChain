package algo

// MergeSortOperation sorts data using the merge sort algorithm.
// It provides stable sorting with O(n log n) time complexity and uses O(n) additional space.
type MergeSortOperation[T any] struct {
	Comparator func(a, b T) bool
}

// Apply performs the merge sort operation on the data.
// It returns a sorted slice based on the provided comparator function.
//
// Example:
//
//	pipeline := NewPipeline[float64]().
//	    MergeSort(func(a, b float64) bool { return a < b })
//	result, err := pipeline.Execute()
func (m *MergeSortOperation[T]) Apply(data []T) ([]T, error) {
	if len(data) < 2 {
		return data, nil
	}

	buffer := make([]T, len(data))
	copy(buffer, data)

	mergeSort(data, buffer, 0, len(data)-1, m.Comparator)

	return data, nil
}
func mergeSort[T any](data, buffer []T, left, right int, cmp func(a, b T) bool) {
	if left < right {
		mid := (left + right) / 2
		mergeSort[T](data, buffer, left, mid, cmp)
		mergeSort[T](data, buffer, mid+1, right, cmp)
		merge(data, buffer, left, mid, right, cmp)
	}
}

// merge merges two sorted slices of data.
func merge[T any](data, buffer []T, left, mid, right int, cmp func(a, b T) bool) {
	copy(buffer[left:right+1], data[left:right+1])

	i := left
	j := mid + 1
	k := left

	for ; k <= right; k++ {
		if i > mid {
			data[k] = buffer[j]
			j++
		} else if j > right {
			data[k] = buffer[i]
			i++
		} else if cmp(buffer[i], buffer[j]) {
			data[k] = buffer[i]
			i++
		} else {
			data[k] = buffer[j]
			j++
		}
	}
}

// MergeSort adds a merge sort operation to the pipeline.
// The comparator function should return true when a should come before b in the sorted result.
//
// Example:
//
//	pipeline.MergeSort(func(a, b Order) bool {
//	    return a.Priority > b.Priority // Sort by priority in descending order
//	})
func (p *Pipeline[T]) MergeSort(comparator func(a, b T) bool) *Pipeline[T] {
	p.operations = append(p.operations, &MergeSortOperation[T]{Comparator: comparator})
	return p
}
