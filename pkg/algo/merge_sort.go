package algo

// MergeSortOperation is an Operation that sorts a slice of data using merge sort.
type MergeSortOperation[T any] struct {
	Comparator func(a, b T) bool
}

// Apply applies the merge sort operation to the data.
func (m *MergeSortOperation[T]) Apply(data []T) ([]T, error) {
	if len(data) < 2 {
		return data, nil
	}

	// 作業用バッファを再利用
	buffer := make([]T, len(data))
	copy(buffer, data)

	// インプレースマージソートの実装
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

// MergeSort sorts a slice of data using merge sort.
func (p *Pipeline[T]) MergeSort(comparator func(a, b T) bool) *Pipeline[T] {
	p.operations = append(p.operations, &MergeSortOperation[T]{Comparator: comparator})
	return p
}
