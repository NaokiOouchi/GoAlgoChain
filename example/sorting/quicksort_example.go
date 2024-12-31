package sorting

import (
	"fmt"

	"github.com/NaokiOouchi/GoAlgoChain/pkg/algo"
)

func QuickSortBasicExample() {
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	result, _ := algo.NewPipelineWithData(numbers).
		QuickSort(func(a, b int) bool { return a < b }).
		Execute()

	fmt.Printf("Sorted numbers: %v\n", result)
}

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

func QuickSortStructExample() {
	products := []Product{
		{ID: 1, Name: "Laptop", Price: 999.99, Stock: 50},
		{ID: 2, Name: "Mouse", Price: 24.99, Stock: 100},
		{ID: 3, Name: "Keyboard", Price: 59.99, Stock: 75},
	}

	// Sort by price ascending
	byPrice, _ := algo.NewPipeline[Product]().
		WithData(products).
		QuickSort(func(a, b Product) bool { return a.Price < b.Price }).
		Execute()

	fmt.Printf("Products sorted by price: %v\n", byPrice)

	// Sort by stock descending
	byStock, _ := algo.NewPipelineWithData(products).
		QuickSort(func(a, b Product) bool { return a.Stock > b.Stock }).
		Execute()

	fmt.Printf("Products sorted by stock (desc): %v\n", byStock)
}

func QuickSortCombinedExample() {
	products := []Product{
		{ID: 1, Name: "Laptop", Price: 999.99, Stock: 50},
		{ID: 2, Name: "Mouse", Price: 24.99, Stock: 0},
		{ID: 3, Name: "Keyboard", Price: 59.99, Stock: 75},
		{ID: 4, Name: "Monitor", Price: 299.99, Stock: 0},
	}

	result, _ := algo.NewPipelineWithData(products).
		Filter(func(p Product) bool { return p.Stock > 0 }).
		QuickSort(func(a, b Product) bool { return a.Price < b.Price }).
		Execute()

	fmt.Printf("In-stock products sorted by price: %v\n", result)
}
