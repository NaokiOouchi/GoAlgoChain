// pkg/algo/binary_search_test.go
package algo

import (
	"sort"
	"strconv"
	"testing"
)

func TestBinarySearchOperation_Found(t *testing.T) {
	pipeline := NewPipeline[Item]().
		BinarySearch(func(a Item) bool { return a.ID == 3 })

	data := []Item{
		{ID: 4, Name: "Item4", Active: false},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 1, Name: "Item1", Active: true},
	}

	sort.Slice(data, func(i, j int) bool { return data[i].ID < data[j].ID })

	pipeline.data = data

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

func TestBinarySearchOperation_StringFound(t *testing.T) {
	pipeline := NewPipeline[string]().
		BinarySearch(func(a string) bool { return a >= "apple" })
	data := []string{"banana", "apple", "orange", "grape"}
	pipeline.data = data
	_, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("Expected target to be found, but got error: %v", err)
	}
	ops := pipeline.GetOperations()
	if len(ops) == 0 {
		t.Fatalf("No operations found in pipeline")
	}
}

func TestBinarySearchOperation_NotFound(t *testing.T) {
	pipeline := NewPipeline[Item]().
		BinarySearch(func(a Item) bool { return a.ID == 5 })

	data := []Item{
		{ID: 4, Name: "Item4", Active: false},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 1, Name: "Item1", Active: true},
	}

	sort.Slice(data, func(i, j int) bool { return data[i].ID < data[j].ID })

	pipeline.data = data

	_, err := pipeline.Execute()
	if err == nil {
		t.Fatalf("Expected error when target is not found, but got nil")
	}

	ops := pipeline.GetOperations()
	if len(ops) == 0 {
		t.Fatalf("No operations found in pipeline")
	}
	bsOp, ok := ops[len(ops)-1].(*BinarySearchOperation[Item])
	if !ok {
		t.Fatalf("Expected last operation to be BinarySearchOperation, but got %T", ops[len(ops)-1])
	}

	if bsOp.GetFoundIndex() != -1 {
		t.Errorf("Expected found index to be -1, but got %d", bsOp.GetFoundIndex())
	}
}

func TestBinarySearchOperation_EmptySlice(t *testing.T) {
	pipeline := NewPipeline[Item]().
		BinarySearch(func(a Item) bool { return a.ID == 1 })

	var data []Item

	pipeline.data = data

	_, err := pipeline.Execute()
	if err == nil {
		t.Fatalf("Expected error when searching in empty slice, but got nil")
	}
}

func BenchmarkBinarySearchOperation(b *testing.B) {
	data := make([]Item, 1000000)
	for i := 0; i < 1000000; i++ {
		data[i] = Item{ID: i, Name: "Item" + strconv.Itoa(i), Active: true}
	}
	pipeline := NewPipeline[Item]().
		BinarySearch(func(a Item) bool { return a.ID == 500000 })
	pipeline.data = data

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := pipeline.Execute()
		if err != nil {
			b.Fatalf("Pipeline execution failed: %v", err)
		}
	}
}
