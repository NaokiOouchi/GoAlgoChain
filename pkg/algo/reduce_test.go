package algo

import (
	"strconv"
	"testing"
)

func TestReduceOperation_SumIDs(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Reduce(func(acc, item Item) Item {
			acc.ID += item.ID
			return acc
		})

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 2, Name: "Item2", Active: false},
		{ID: 3, Name: "Item3", Active: true},
	}

	expected := []Item{
		{ID: 6, Name: "Item1", Active: true}, // NameとActiveはaccから継承
	}

	pipeline.WithData(data)

	reducedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("ReduceOperation failed: %v", err)
	}

	if len(reducedData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(reducedData))
	}

	for i, item := range reducedData {
		if item.ID != expected[i].ID {
			t.Errorf("At index %d, expected ID %d, got %d", i, expected[i].ID, item.ID)
		}
		// 他のフィールドはリデュースの目的によっては無視
	}
}

func TestReduceOperation_MaxID(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Reduce(func(acc, item Item) Item {
			if item.ID > acc.ID {
				return item
			}
			return acc
		})

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
		{ID: 3, Name: "Item3", Active: true},
		{ID: 2, Name: "Item2", Active: false},
	}

	expected := []Item{
		{ID: 3, Name: "Item3", Active: true},
	}

	pipeline.WithData(data)

	reducedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("ReduceOperation failed: %v", err)
	}

	if len(reducedData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(reducedData))
	}

	if reducedData[0] != expected[0] {
		t.Errorf("Expected %+v, got %+v", expected[0], reducedData[0])
	}
}

func TestReduceOperation_SingleElement(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Reduce(func(acc, item Item) Item {
			acc.Name = "Reduced" + acc.Name
			return acc
		})

	data := []Item{
		{ID: 1, Name: "Item1", Active: true},
	}

	expected := []Item{
		{ID: 1, Name: "Item1", Active: true},
	}

	pipeline.WithData(data)

	reducedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("ReduceOperation failed: %v", err)
	}

	if len(reducedData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(reducedData))
	}

	if reducedData[0].Name != expected[0].Name {
		t.Errorf("Expected Name %s, got %s", expected[0].Name, reducedData[0].Name)
	}
}

func TestReduceOperation_EmptySlice(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Reduce(func(acc, item Item) Item {
			return acc
		})

	var data []Item

	pipeline.WithData(data)

	_, err := pipeline.Execute()
	if err == nil {
		t.Fatalf("Expected error when reducing empty slice, but got nil")
	}

	expectedError := "ReduceOperation: cannot reduce an empty slice"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}
}

func TestReduceOperation_AllSame(t *testing.T) {
	pipeline := NewPipeline[Item]().
		Reduce(func(acc, item Item) Item {
			return acc
		})

	data := []Item{
		{ID: 2, Name: "Item2", Active: true},
		{ID: 2, Name: "Item2a", Active: false},
		{ID: 2, Name: "Item2b", Active: true},
	}

	expected := []Item{
		{ID: 2, Name: "Item2", Active: true},
	}

	pipeline.WithData(data)

	reducedData, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("ReduceOperation failed: %v", err)
	}

	if len(reducedData) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(reducedData))
	}

	if reducedData[0] != expected[0] {
		t.Errorf("Expected %+v, got %+v", expected[0], reducedData[0])
	}
}

func BenchmarkReduce(b *testing.B) {
	pipeline := NewPipeline[Item]().
		Reduce(func(acc, item Item) Item {
			return acc
		})
	data := make([]Item, 1000000)
	for i := 0; i < 1000000; i++ {
		data[i] = Item{ID: i, Name: "Item" + strconv.Itoa(i), Active: i%2 == 0}
	}
	pipeline.WithData(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := pipeline.Execute()
		if err != nil {
			b.Fatalf("ReduceOperation failed: %v", err)
		}
	}
}
