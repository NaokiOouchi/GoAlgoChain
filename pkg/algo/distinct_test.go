package algo

import (
	"strconv"
	"testing"
)

func TestDistinctOperation_RemoveDuplicates(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Distinct(func(a, b Item) bool { return a.ID == b.ID })

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 1, Name: "Item1a", Active: true},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 2, Name: "Item2a", Active: false},
	}

	expected := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
	}

	pipeline.WithData(data)

	distinctData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("DistinctOperation failed: %v", err)
	}

	if len(distinctData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(distinctData))
	}

	for i, item := range distinctData {
		if item.ID != expected[i].ID || item.Name != expected[i].Name || item.Active != expected[i].Active {
			t.Errorf("At index %d, expected %+v, got %+v", i, expected[i], item)
		}
	}
}

func TestDistinctOperation_AllUnique(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Distinct(func(a, b Item) bool { return a.ID == b.ID })

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
	}

	expected := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
	}

	pipeline.WithData(data)

	distinctData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("DistinctOperation failed: %v", err)
	}

	if len(distinctData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(distinctData))
	}

	for i, item := range distinctData {
		if item != expected[i] {
			t.Errorf("At index %d, expected %+v, got %+v", i, expected[i], item)
		}
	}
}

func TestDistinctOperation_AllDuplicates(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Distinct(func(a, b Item) bool { return a.ID == b.ID })

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 1, Name: "Item1a", Active: false},
		{ID: 1, Name: "Item1b", Active: true},
	}

	expected := []Item{
		{ID: 1, Name: "Item1", Active: true},
	}

	pipeline.WithData(data)

	distinctData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("DistinctOperation failed: %v", err)
	}

	if len(distinctData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(distinctData))
	}

	if distinctData[0] != expected[0] {
		t.Errorf("Expected %+v, got %+v", expected[0], distinctData[0])
	}
}

func TestDistinctOperation_EmptySlice(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Distinct(func(a, b Item) bool { return a.ID == b.ID })

	var data []Item

	pipeline.WithData(data)

	distinctData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("DistinctOperation failed on empty slice: %v", err)
	}

	if len(distinctData) != 0 {
		t.Errorf("Expected 0 items, but got %+v", distinctData)
	}
}

func TestDistinctOperation_MultipleCriteria(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Distinct(func(a, b Item) bool { return a.ID == b.ID && a.Active == b.Active })

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 1, Name: "Item1a", Active: true},
		{ID: 1, Name: "Item1b", Active: false},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 2, Name: "Item2a", Active: false},
	}

	expected := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 1, Name: "Item1b", Active: false},
		{ID: 2, Name: "Item2", Active: false},
	}

	pipeline.WithData(data)

	distinctData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("DistinctOperation failed: %v", err)
	}

	if len(distinctData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(distinctData))
	}

	for i, item := range distinctData {
		if item != expected[i] {
			t.Errorf("At index %d, expected %+v, got %+v", i, expected[i], item)
		}
	}
}

func BenchmarkDistinct(b *testing.B) {
	pipeline := NewPipeline[Item]().
		Distinct(func(a, b Item) bool { return a.ID == b.ID })
	data := make([]Item, 100000)
	for i := 0; i < 100000; i++ {
		data[i] = Item{ID: i, Name: "Item" + strconv.Itoa(i), Active: true}
	}
	pipeline.WithData(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pipeline.Execute()
	}
}
