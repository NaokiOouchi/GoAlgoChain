package algo

import (
	"testing"
)

func TestLinearSearchOperation_ExactMatch(t *testing.T) {
	pipeline := NewPipeline[Item]().
		LinearSearchExact(Item{ID: 3, Name: "Item3", Active: true})

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 4, Name: "Item4", Active: false},
	}

	pipeline.WithData(data)

	_, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("Expected target to be found, but got error: %v", err)
	}
}

func TestLinearSearchOperation_PartialMatch(t *testing.T) {
	pipeline := NewPipeline[Item]().
		LinearSearch(func(a Item) bool {
			return a.ID == 3
		})

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 4, Name: "Item4", Active: false},
	}

	pipeline.WithData(data)

	_, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("Expected target to be found, but got error: %v", err)
	}
}

func TestLinearSearchOperation_NotFound(t *testing.T) {
	pipeline := NewPipeline[Item]().
		LinearSearchExact(Item{ID: 5, Name: "Item5", Active: true})

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 4, Name: "Item4", Active: false},
	}

	pipeline.WithData(data)

	_, err := pipeline.Execute()
	if err == nil {
		t.Fatalf("Expected error when target is not found, but got nil")
	}
}

func TestLinearSearchOperation_EmptySlice(t *testing.T) {
	pipeline := NewPipeline[int]().
		LinearSearch(func(item int) bool {
			return item == 1
		})

	var data []int

	pipeline.WithData(data)

	_, err := pipeline.Execute()
	if err == nil {
		t.Fatalf("Expected error when searching in empty slice, but got nil")
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	pipeline := NewPipeline[int]().
		LinearSearch(func(item int) bool {
			return item == 1
		})
	data := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		data[i] = i
	}
	pipeline.WithData(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := pipeline.Execute()
		if err != nil {
			b.Fatalf("Pipeline execution failed: %v", err)
		}
	}
}
