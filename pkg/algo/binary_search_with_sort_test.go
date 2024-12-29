// pkg/algo/binary_search_with_sort_test.go
package algo

import (
	"testing"
)

func TestBinarySearchOperation_WithQuickSort_Found(t *testing.T) {
	pipeline := NewPipeline[Item]().
		QuickSort(func(a, b Item) bool { return a.ID < b.ID }).
		BinarySearch(func(a Item) bool { return a.ID == 3 })

	data := []Item{
		{ID: 4, Name: "Item4", Active: false},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 1, Name: "Item1", Active: true},
	}
	pipeline.WithData(data)
	_, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("Expected target to be found, but got error: %v", err)
	}
	ops := pipeline.GetOperations()
	if len(ops) == 0 {
		t.Fatalf("No operations found in pipeline")
	}
	bsOp, ok := ops[len(ops)-1].(*BinarySearchOperation[Item])
	if !ok {
		t.Fatalf("Expected last operation to be BinarySearchOperation, but got %T", ops[len(ops)-1])
	}
	expectedIndex := 2
	if bsOp.GetFoundIndex() != expectedIndex {
		t.Errorf("Expected found index to be %d, but got %d", expectedIndex, bsOp.GetFoundIndex())
	}
}

func TestBinarySearchOperation_String_Found(t *testing.T) {
	pipeline := NewPipeline[string]().
		QuickSort(func(a, b string) bool { return a < b }).
		BinarySearch(func(a string) bool { return a >= "apple" })

	data := []string{"banana", "apple", "orange", "grape"}
	pipeline.WithData(data)

	_, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("Expected target to be found, but got error: %v", err)
	}

	ops := pipeline.GetOperations()
	if len(ops) == 0 {
		t.Fatalf("No operations found in pipeline")
	}

	bsOp, ok := ops[len(ops)-1].(*BinarySearchOperation[string])
	if !ok {
		t.Fatalf("Expected last operation to be BinarySearchOperation, but got %T", ops[len(ops)-1])
	}

	expectedIndex := 0
	if bsOp.GetFoundIndex() != expectedIndex {
		t.Errorf("Expected found index to be %d, but got %d", expectedIndex, bsOp.GetFoundIndex())
	}
}

func TestBinarySearchOperation_WithQuickSort_NotFound(t *testing.T) {
	pipeline := NewPipeline[Item]().
		QuickSort(func(a, b Item) bool { return a.ID < b.ID }).
		BinarySearch(func(a Item) bool { return a.ID == 5 })
	data := []Item{
		{ID: 4, Name: "Item4", Active: false},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 1, Name: "Item1", Active: true},
	}
	pipeline.WithData(data)
	_, err := pipeline.Execute()
	if err == nil {
		t.Fatalf("Expected error when target is not found after sorting, but got nil")
	}
	expectedError := "target not found in data"
	if err.Error() != expectedError {
		t.Errorf("Expected error message to be '%s', but got '%s'", expectedError, err.Error())
	}
}

func TestBinarySearchOperation_WithQuickSort_EmptySlice(t *testing.T) {
	pipeline := NewPipeline[Item]().
		QuickSort(func(a, b Item) bool { return a.ID < b.ID }).
		BinarySearch(func(a Item) bool { return a.ID == 3 })
	var data []Item
	pipeline.WithData(data)
	_, err := pipeline.Execute()
	if err == nil {
		t.Fatalf("Expected error when searching in empty slice, but got nil")
	}
}
