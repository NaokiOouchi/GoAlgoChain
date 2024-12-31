package filtering

import (
	"fmt"

	"github.com/NaokiOouchi/GoAlgoChain/pkg/algo"
)

func FilterBasicExample() {
	numbers := []int{1, -2, 3, -4, 5, -6, 7, -8}
	result, _ := algo.NewPipelineWithData(numbers).
		Filter(func(n int) bool { return n > 0 }).
		Execute()

	fmt.Printf("Positive numbers: %v\n", result)
}

type Transaction struct {
	ID     int
	Amount float64
	Status string
	UserID int
}

func FilterStructExample() {
	transactions := []Transaction{
		{ID: 1, Amount: 100.50, Status: "completed", UserID: 1},
		{ID: 2, Amount: 25.99, Status: "pending", UserID: 2},
		{ID: 3, Amount: 75.00, Status: "completed", UserID: 1},
		{ID: 4, Amount: 50.25, Status: "failed", UserID: 3},
	}

	// Filter completed transactions
	completed, _ := algo.NewPipelineWithData(transactions).
		Filter(func(t Transaction) bool { return t.Status == "completed" }).
		Execute()

	fmt.Printf("Completed transactions: %v\n", completed)
}

func FilterCombinedExample() {
	transactions := []Transaction{
		{ID: 1, Amount: 100.50, Status: "completed", UserID: 1},
		{ID: 2, Amount: 25.99, Status: "completed", UserID: 2},
		{ID: 3, Amount: 75.00, Status: "completed", UserID: 1},
	}

	result, _ := algo.NewPipelineWithData(transactions).
		Filter(func(t Transaction) bool { return t.UserID == 1 }).
		Filter(func(t Transaction) bool { return t.Amount > 50 }).
		Execute()

	fmt.Printf("User 1's large transactions: %v\n", result)
}
