package algo

import (
	"testing"
)

func TestFilterOperation_FilterActiveItems(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Filter(func(a Item) bool { return a.Active })

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 4, Name: "Item4", Active: false},
	}

	expected := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 3, Name: "Item3", Active: true},
	}

	pipeline.data = data

	sortedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("FilterOperation failed: %v", err)
	}

	if len(sortedData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(sortedData))
	}

	for i, item := range sortedData {
		if item != expected[i] {
			t.Errorf("At index %d, expected %+v, got %+v", i, expected[i], item)
		}
	}
}

func TestFilterOperation_FilterByID(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Filter(func(a Item) bool { return a.ID == 3 })

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 3, Name: "Item3b", Active: true},
	}

	expected := []Item{
		{ID: 3, Name: "Item3", Active: true},
		{ID: 3, Name: "Item3b", Active: true},
	}

	pipeline.data = data

	sortedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("FilterOperation failed: %v", err)
	}

	if len(sortedData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(sortedData))
	}

	for i, item := range sortedData {
		if item != expected[i] {
			t.Errorf("At index %d, expected %+v, got %+v", i, expected[i], item)
		}
	}
}

func TestFilterOperation_NoMatch(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Filter(func(a Item) bool { return a.ID == 5 })

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 4, Name: "Item4", Active: false},
	}

	pipeline.data = data

	sortedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("FilterOperation failed: %v", err)
	}

	if len(sortedData) != 0 {
		t.Errorf("Expected no items, but got %+v", sortedData)
	}
}

func TestFilterOperation_EmptySlice(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Filter(func(a Item) bool { return a.Active })

	var data []Item

	pipeline.data = data

	sortedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("FilterOperation failed on empty slice: %v", err)
	}

	if len(sortedData) != 0 {
		t.Errorf("Expected no items, but got %+v", sortedData)
	}
}

// TestFilterOperation_AllMatch は、すべてのアイテムがフィルタ条件を満たす場合のテストケースです。
func TestFilterOperation_AllMatch(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Filter(func(a Item) bool { return a.Active })

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 5, Name: "Item5", Active: true},
	}

	expected := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 5, Name: "Item5", Active: true},
	}

	pipeline.data = data

	sortedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("FilterOperation failed: %v", err)
	}

	if len(sortedData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(sortedData))
	}

	for i, item := range sortedData {
		if item != expected[i] {
			t.Errorf("At index %d, expected %+v, got %+v", i, expected[i], item)
		}
	}
}

// BenchmarkFilterOperation benchmarks the FilterOperation.
func BenchmarkFilterOperation(b *testing.B) {
	data := make([]int, 1000000)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	predicate := func(x int) bool {
		return x%2 == 0
	}

	pipeline := NewPipeline[int]().
		Filter(predicate)
	pipeline.data = data

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := pipeline.Execute()
		if err != nil {
			b.Fatalf("Pipeline execution failed: %v", err)
		}
	}
}
