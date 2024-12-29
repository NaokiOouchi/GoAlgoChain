package algo

import (
	"reflect"
	"testing"
)

func TestGroupBy_SimpleGrouping(t *testing.T) {
	data := []User{
		{ID: 1, Name: "Alice", Active: true},
		{ID: 2, Name: "Bob", Active: false},
		{ID: 1, Name: "Alice A", Active: true},
		{ID: 3, Name: "Charlie", Active: true},
		{ID: 2, Name: "Bob B", Active: false},
	}

	expected := []GroupedItem[int, User]{
		{
			Key: 1,
			Items: []User{
				{ID: 1, Name: "Alice", Active: true},
				{ID: 1, Name: "Alice A", Active: true},
			},
		},
		{
			Key: 2,
			Items: []User{
				{ID: 2, Name: "Bob", Active: false},
				{ID: 2, Name: "Bob B", Active: false},
			},
		},
		{
			Key: 3,
			Items: []User{
				{ID: 3, Name: "Charlie", Active: true},
			},
		},
	}

	result := GroupBy(data, func(u User) int { return u.ID })

	if !compareGroupedItems(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGroupBy_EmptySlice(t *testing.T) {
	var data []User

	var expected []GroupedItem[int, User]

	result := GroupBy(data, func(u User) int { return u.ID })

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected empty slice, got %v", result)
	}
}

func TestGroupBy_AllSameKey(t *testing.T) {
	data := []Order{
		{OrderID: 101, UserID: 1, Item: "Book"},
		{OrderID: 102, UserID: 1, Item: "Pen"},
		{OrderID: 103, UserID: 1, Item: "Notebook"},
	}

	expected := []GroupedItem[int, Order]{
		{
			Key: 1,
			Items: []Order{
				{OrderID: 101, UserID: 1, Item: "Book"},
				{OrderID: 102, UserID: 1, Item: "Pen"},
				{OrderID: 103, UserID: 1, Item: "Notebook"},
			},
		},
	}

	result := GroupBy(data, func(o Order) int { return o.UserID })

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGroupBy_AllUniqueKeys(t *testing.T) {
	data := []User{
		{ID: 1, Name: "Alice", Active: true},
		{ID: 2, Name: "Bob", Active: false},
		{ID: 3, Name: "Charlie", Active: true},
	}

	expected := []GroupedItem[int, User]{
		{
			Key: 1,
			Items: []User{
				{ID: 1, Name: "Alice", Active: true},
			},
		},
		{
			Key: 2,
			Items: []User{
				{ID: 2, Name: "Bob", Active: false},
			},
		},
		{
			Key: 3,
			Items: []User{
				{ID: 3, Name: "Charlie", Active: true},
			},
		},
	}

	result := GroupBy(data, func(u User) int { return u.ID })

	if !compareGroupedItems(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGroupBy_CustomKeyFunction(t *testing.T) {
	data := []User{
		{ID: 1, Name: "Alice", Active: true},
		{ID: 2, Name: "Bob", Active: false},
		{ID: 3, Name: "Charlie", Active: true},
		{ID: 4, Name: "David", Active: false},
	}

	expected := []GroupedItem[bool, User]{
		{
			Key: true,
			Items: []User{
				{ID: 1, Name: "Alice", Active: true},
				{ID: 3, Name: "Charlie", Active: true},
			},
		},
		{
			Key: false,
			Items: []User{
				{ID: 2, Name: "Bob", Active: false},
				{ID: 4, Name: "David", Active: false},
			},
		},
	}

	result := GroupBy(data, func(u User) bool { return u.Active })

	if !compareGroupedItems(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGroupBy_StructKey(t *testing.T) {
	type Product struct {
		Category string
		Name     string
	}

	data := []Product{
		{Category: "Electronics", Name: "Laptop"},
		{Category: "Electronics", Name: "Smartphone"},
		{Category: "Furniture", Name: "Chair"},
		{Category: "Furniture", Name: "Table"},
		{Category: "Electronics", Name: "Tablet"},
	}

	expected := []GroupedItem[string, Product]{
		{
			Key: "Electronics",
			Items: []Product{
				{Category: "Electronics", Name: "Laptop"},
				{Category: "Electronics", Name: "Smartphone"},
				{Category: "Electronics", Name: "Tablet"},
			},
		},
		{
			Key: "Furniture",
			Items: []Product{
				{Category: "Furniture", Name: "Chair"},
				{Category: "Furniture", Name: "Table"},
			},
		},
	}

	result := GroupBy(data, func(p Product) string { return p.Category })

	if !compareGroupedItems(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Helper function to compare two slices of GroupedItem
func compareGroupedItems[T any, K comparable](a, b []GroupedItem[K, T]) bool {
	if len(a) != len(b) {
		return false
	}

	aMap := make(map[K][]T)
	for _, group := range a {
		aMap[group.Key] = group.Items
	}

	for _, group := range b {
		items, exists := aMap[group.Key]
		if !exists {
			return false
		}
		if !reflect.DeepEqual(items, group.Items) {
			return false
		}
	}

	return true
}
