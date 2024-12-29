package algo

import (
	"testing"
)

func TestMapOperation_IncrementID(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Map(func(a Item) Item {
			a.ID += 1
			return a
		})

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
	}

	expected := []Item{
		{ID: 2, Name: "Item1", Active: true},
		{ID: 3, Name: "Item2", Active: false},
		{ID: 4, Name: "Item3", Active: true},
	}

	pipeline.data = data

	mappedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("MapOperation failed: %v", err)
	}

	if len(mappedData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(mappedData))
	}

	for i, item := range mappedData {
		if item != expected[i] {
			t.Errorf("At index %d, expected %+v, got %+v", i, expected[i], item)
		}
	}
}

func TestMapOperation_Identity(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Map(func(a Item) Item {
			return a
		})

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
	}

	expected := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
	}

	pipeline.data = data

	mappedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("MapOperation failed: %v", err)
	}

	if len(mappedData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(mappedData))
	}

	for i, item := range mappedData {
		if item != expected[i] {
			t.Errorf("At index %d, expected %+v, got %+v", i, expected[i], item)
		}
	}
}

func TestMapOperation_EmptySlice(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Map(func(a Item) Item {
			a.ID += 1
			return a
		})

	var data []Item

	pipeline.data = data

	mappedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("MapOperation failed on empty slice: %v", err)
	}

	if len(mappedData) != 0 {
		t.Errorf("Expected 0 items, but got %d", len(mappedData))
	}
}

func TestMapOperation_AllMapped(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Map(func(a Item) Item {
			a.Name = "Mapped" + a.Name
			return a
		})

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
	}

	expected := []Item{
		{ID: 1, Name: "MappedItem1", Active: true},
		{ID: 2, Name: "MappedItem2", Active: false},
	}

	pipeline.data = data

	mappedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("MapOperation failed: %v", err)
	}

	if len(mappedData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(mappedData))
	}

	for i, item := range mappedData {
		if item != expected[i] {
			t.Errorf("At index %d, expected %+v, got %+v", i, expected[i], item)
		}
	}
}
