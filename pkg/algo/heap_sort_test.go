package algo

import (
	"reflect"
	"testing"
)

func TestHeapSortOperation(t *testing.T) {
	pipeline := NewPipeline[Item]().
		HeapSort(func(a, b Item) int { return a.ID - b.ID })

	data := []Item{
		{ID: 3, Name: "Item3", Active: true},
		{ID: 1, Name: "Item1", Active: true},
		{ID: 4, Name: "Item4", Active: false},
		{ID: 2, Name: "Item2", Active: false},
	}

	pipeline.data = data

	sortedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	expected := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 4, Name: "Item4", Active: false},
	}

	if !reflect.DeepEqual(sortedData, expected) {
		t.Errorf("Expected %v, got %v", expected, sortedData)
	}
}

func TestHeapSortOperation_AlreadySorted(t *testing.T) {
	pipeline := NewPipeline[int]().
		HeapSort(func(a, b int) int { return a - b })

	data := []int{1, 2, 3, 4, 5}

	pipeline.data = data

	sortedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	expected := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(sortedData, expected) {
		t.Errorf("Expected %v, got %v", expected, sortedData)
	}
}

func TestHeapSortOperation_EmptySlice(t *testing.T) {
	pipeline := NewPipeline[string]().
		HeapSort(func(a, b string) int { return len(a) - len(b) })

	var data []string

	pipeline.data = data

	sortedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	var expected []string

	if !reflect.DeepEqual(sortedData, expected) {
		t.Errorf("Expected empty slice, got %v", sortedData)
	}
}
