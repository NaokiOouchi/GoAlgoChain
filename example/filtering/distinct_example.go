package filtering

import (
	"fmt"

	"github.com/NaokiOouchi/GoAlgoChain/pkg/algo"
)

func DistinctBasicExample() {
	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	result, _ := algo.NewPipelineWithData(numbers).
		Distinct(func(a, b int) bool { return a == b }).
		Execute()

	fmt.Printf("Unique numbers: %v\n", result)
}

type Employee struct {
	ID         int
	Department string
	Role       string
}

func DistinctStructExample() {
	employees := []Employee{
		{ID: 1, Department: "Engineering", Role: "Developer"},
		{ID: 2, Department: "Sales", Role: "Manager"},
		{ID: 3, Department: "Engineering", Role: "Developer"},
		{ID: 4, Department: "Marketing", Role: "Manager"},
	}

	// Distinct by department
	byDepartment, _ := algo.NewPipelineWithData(employees).
		Distinct(func(a, b Employee) bool { return a.Department == b.Department }).
		Execute()

	fmt.Printf("Unique departments: %v\n", byDepartment)
}

func DistinctCombinedExample() {
	employees := []Employee{
		{ID: 1, Department: "Engineering", Role: "Developer"},
		{ID: 2, Department: "Sales", Role: "Manager"},
		{ID: 3, Department: "Engineering", Role: "Developer"},
	}

	result, _ := algo.NewPipelineWithData(employees).
		Filter(func(e Employee) bool { return e.Department == "Engineering" }).
		Distinct(func(a, b Employee) bool { return a.Role == b.Role }).
		Execute()

	fmt.Printf("Unique roles in Engineering: %v\n", result)
}
