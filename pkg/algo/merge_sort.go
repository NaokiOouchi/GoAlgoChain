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
	mid := len(data) / 2
	left, err := (&MergeSortOperation[T]{Comparator: m.Comparator}).Apply(data[:mid])
	if err != nil {
		return nil, err
	}
	right, err := (&MergeSortOperation[T]{Comparator: m.Comparator}).Apply(data[mid:])
	if err != nil {
		return nil, err
	}
	return merge(left, right, m.Comparator), nil
}

// merge merges two sorted slices of data.
func merge[T any](left, right []T, cmp func(a, b T) bool) []T {
	result := make([]T, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if cmp(left[i], right[j]) {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// MergeSort sorts a slice of data using merge sort.
func (p *Pipeline[T]) MergeSort(comparator func(a, b T) bool) *Pipeline[T] {
	p.operations = append(p.operations, &MergeSortOperation[T]{Comparator: comparator})
	return p
}
