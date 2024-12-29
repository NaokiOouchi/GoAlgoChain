package algo

import (
	"testing"
)

func TestSkipOperation_SkipN(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Skip(2)

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 4, Name: "Item4", Active: false},
	}

	expected := []Item{
		{ID: 3, Name: "Item3", Active: true},
		{ID: 4, Name: "Item4", Active: false},
	}

	pipeline.data = data

	skippedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("SkipOperation failed: %v", err)
	}

	if len(skippedData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(skippedData))
	}

	for i, item := range skippedData {
		if item != expected[i] {
			t.Errorf("At index %d, expected %+v, got %+v", i, expected[i], item)
		}
	}
}

func TestSkipOperation_SkipMoreThanLength(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Skip(5)

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
	}

	pipeline.data = data

	skippedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("SkipOperation failed: %v", err)
	}

	if len(skippedData) != 0 {
		t.Errorf("Expected 0 items, but got %+v", skippedData)
	}
}

func TestSkipOperation_SkipZero(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Skip(0)

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
	}

	expected := []Item{
		{ID: 1, Name: "Item1", Active: true},
	}

	pipeline.data = data

	skippedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("SkipOperation failed: %v", err)
	}

	if len(skippedData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(skippedData))
	}

	for i, item := range skippedData {
		if item != expected[i] {
			t.Errorf("At index %d, expected %+v, got %+v", i, expected[i], item)
		}
	}
}

func TestSkipOperation_EmptySlice(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Skip(3)

	var data []Item

	pipeline.data = data

	skippedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("SkipOperation failed on empty slice: %v", err)
	}

	if len(skippedData) != 0 {
		t.Errorf("Expected 0 items, but got %+v", skippedData)
	}
}
