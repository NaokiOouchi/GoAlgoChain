package combined

import (
	"fmt"

	"github.com/NaokiOouchi/GoAlgoChain/pkg/algo"
)

type Order struct {
	ID       int
	UserID   int
	Amount   float64
	Status   string
	Category string
}

func ComplexPipelineExample() {
	orders := []Order{
		{ID: 1, UserID: 1, Amount: 100.50, Status: "completed", Category: "Electronics"},
		{ID: 2, UserID: 2, Amount: 25.99, Status: "pending", Category: "Books"},
		{ID: 3, UserID: 1, Amount: 75.00, Status: "completed", Category: "Electronics"},
		{ID: 4, UserID: 3, Amount: 50.25, Status: "completed", Category: "Books"},
		{ID: 5, UserID: 2, Amount: 30.00, Status: "completed", Category: "Electronics"},
	}

	result, _ := algo.NewPipelineWithData(orders).
		Filter(func(o Order) bool { return o.Status == "completed" }).
		Distinct(func(a, b Order) bool { return a.Category == b.Category }).
		QuickSort(func(a, b Order) bool { return a.Amount > b.Amount }).
		Take(2).
		Execute()

	fmt.Printf("Top 2 categories by highest order amount: %v\n", result)

	// Group the results by category
	grouped := algo.GroupBy(result, func(o Order) string {
		return o.Category
	})

	fmt.Printf("Grouped by category: %v\n", grouped)
}
