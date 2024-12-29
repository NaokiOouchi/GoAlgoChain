package algo

import (
	"reflect"
	"testing"
)

func TestMergeSortOperation(t *testing.T) {
	pipeline := NewPipeline[Item]().
		MergeSort(func(a, b Item) bool { return a.ID < b.ID })

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

func TestMergeSortOperation_AlreadySorted(t *testing.T) {
	pipeline := NewPipeline[int]().
		MergeSort(func(a, b int) bool { return a < b })

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

func TestMergeSortOperation_EmptySlice(t *testing.T) {
	pipeline := NewPipeline[string]().
		MergeSort(func(a, b string) bool { return a < b })

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
