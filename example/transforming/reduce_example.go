package transforming

import (
	"fmt"

	"github.com/NaokiOouchi/GoAlgoChain/pkg/algo"
)

func ReduceBasicExample() {
	numbers := []int{1, 2, 3, 4, 5}
	result, _ := algo.NewPipelineWithData(numbers).
		Reduce(func(acc, item int) int { return acc + item }).
		Execute()

	fmt.Printf("Sum of numbers: %v\n", result)
}

type Order struct {
	ID     int
	Amount float64
	UserID int
}

func ReduceStructExample() {
	orders := []Order{
		{ID: 1, Amount: 100.50, UserID: 1},
		{ID: 2, Amount: 25.99, UserID: 1},
		{ID: 3, Amount: 75.00, UserID: 2},
	}

	// Calculate total amount
	result, _ := algo.NewPipelineWithData(orders).
		Reduce(func(acc, item Order) Order {
			acc.Amount += item.Amount
			return acc
		}).
		Execute()

	fmt.Printf("Total order amount: %.2f\n", result[0].Amount)
}

func ReduceCombinedExample() {
	orders := []Order{
		{ID: 1, Amount: 100.50, UserID: 1},
		{ID: 2, Amount: 25.99, UserID: 2},
		{ID: 3, Amount: 75.00, UserID: 1},
	}

	result, _ := algo.NewPipelineWithData(orders).
		Filter(func(o Order) bool { return o.UserID == 1 }).
		Reduce(func(acc, item Order) Order {
			acc.Amount += item.Amount
			return acc
		}).
		Execute()

	fmt.Printf("Total amount for User 1: %.2f\n", result[0].Amount)
}
