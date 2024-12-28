package algo

import (
	"reflect"
	"testing"
)

// Item is a test struct for the FindOperation.
type Item struct {
	ID     int
	Name   string
	Active bool
}

func TestFindOperation(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Find(func(item Item) bool { return item.Active })

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 4, Name: "Item4", Active: false},
	}

	pipeline.data = data

	result, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	expected := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 3, Name: "Item3", Active: true},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFindOperation_NoMatches(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Find(func(item Item) bool { return item.ID > 100 })

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 4, Name: "Item4", Active: false},
	}

	pipeline.data = data

	result, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	var expected []Item

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected empty slice, got %v", result)
	}
}
