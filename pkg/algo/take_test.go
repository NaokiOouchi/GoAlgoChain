package algo

import (
	"strconv"
	"testing"
)

func TestTakeOperation_TakeN(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Take(2)

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
	}

	expected := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
	}

	pipeline.data = data

	takenData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("TakeOperation failed: %v", err)
	}

	if len(takenData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(takenData))
	}

	for i, item := range takenData {
		if item != expected[i] {
			t.Errorf("At index %d, expected %+v, got %+v", i, expected[i], item)
		}
	}
}

func TestTakeOperation_TakeMoreThanLength(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Take(5)

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
	}

	expected := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
	}

	pipeline.data = data

	takenData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("TakeOperation failed: %v", err)
	}

	if len(takenData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(takenData))
	}

	for i, item := range takenData {
		if item != expected[i] {
			t.Errorf("At index %d, expected %+v, got %+v", i, expected[i], item)
		}
	}
}

func TestTakeOperation_TakeZero(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Take(0)

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
	}

	pipeline.data = data

	takenData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("TakeOperation failed: %v", err)
	}

	if len(takenData) != 0 {
		t.Errorf("Expected 0 items, but got %+v", takenData)
	}
}

func TestTakeOperation_TakeNegative(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Take(-3)

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
	}

	pipeline.data = data

	takenData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("TakeOperation failed: %v", err)
	}

	if len(takenData) != 0 {
		t.Errorf("Expected 0 items, but got %+v", takenData)
	}
}

func BenchmarkTake(b *testing.B) {
	pipeline := NewPipeline[Item]().
		Take(100)
	data := make([]Item, 1000000)
	for i := 0; i < 1000000; i++ {
		data[i] = Item{ID: i, Name: "Item" + strconv.Itoa(i), Active: true}
	}
	pipeline.data = data
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := pipeline.Execute()
		if err != nil {
			b.Fatalf("TakeOperation failed: %v", err)
		}
	}
}
