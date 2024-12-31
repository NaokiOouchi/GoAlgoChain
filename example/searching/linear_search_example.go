package searching

import (
	"fmt"

	"github.com/NaokiOouchi/GoAlgoChain/pkg/algo"
)

func LinearSearchBasicExample() {
	numbers := []int{45, 12, 78, 23, 7, 19, 3}
	result, _ := algo.NewPipelineWithData(numbers).
		LinearSearch(func(n int) bool { return n == 23 }).
		Execute()

	fmt.Printf("Found number in unsorted array: %v\n", result)
}

type Product struct {
	ID      int
	Name    string
	Price   float64
	InStock bool
}

func LinearSearchExactExample() {
	target := Product{ID: 2, Name: "Mouse", Price: 24.99, InStock: true}
	products := []Product{
		{ID: 1, Name: "Keyboard", Price: 59.99, InStock: true},
		target,
		{ID: 3, Name: "Monitor", Price: 299.99, InStock: false},
	}

	result, _ := algo.NewPipelineWithData(products).
		LinearSearchExact(target).
		Execute()

	fmt.Printf("Found exact product match: %v\n", result)
}

func LinearSearchCombinedExample() {
	products := []Product{
		{ID: 1, Name: "Keyboard", Price: 59.99, InStock: true},
		{ID: 2, Name: "Mouse", Price: 24.99, InStock: false},
		{ID: 3, Name: "Monitor", Price: 299.99, InStock: true},
	}

	result, _ := algo.NewPipelineWithData(products).
		Filter(func(p Product) bool { return p.InStock }).
		LinearSearch(func(p Product) bool { return p.Price > 100 }).
		Execute()

	fmt.Printf("Found in-stock product over $100: %v\n", result)
}
