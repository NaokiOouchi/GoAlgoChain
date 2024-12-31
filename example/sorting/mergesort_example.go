package sorting

import (
	"fmt"

	"github.com/NaokiOouchi/GoAlgoChain/pkg/algo"
)

func MergeSortBasicExample() {
	numbers := []int{38, 27, 43, 3, 9, 82, 10}
	result, _ := algo.NewPipelineWithData(numbers).
		MergeSort(func(a, b int) bool { return a < b }).
		Execute()

	fmt.Printf("Sorted numbers: %v\n", result)
}

type Student struct {
	ID     int
	Name   string
	Grade  float64
	Active bool
}

func MergeSortStructExample() {
	students := []Student{
		{ID: 1, Name: "Alice", Grade: 3.8, Active: true},
		{ID: 2, Name: "Bob", Grade: 3.5, Active: true},
		{ID: 3, Name: "Charlie", Grade: 4.0, Active: false},
	}

	// Sort by grade descending
	byGrade, _ := algo.NewPipeline[Student]().
		WithData(students).
		MergeSort(func(a, b Student) bool { return a.Grade > b.Grade }).
		Execute()

	fmt.Printf("Students sorted by grade: %v\n", byGrade)
}

func MergeSortCombinedExample() {
	students := []Student{
		{ID: 1, Name: "Alice", Grade: 3.8, Active: true},
		{ID: 2, Name: "Bob", Grade: 3.5, Active: false},
		{ID: 3, Name: "Charlie", Grade: 4.0, Active: true},
	}

	result, _ := algo.NewPipelineWithData(students).
		Filter(func(s Student) bool { return s.Active }).
		MergeSort(func(a, b Student) bool { return a.Grade > b.Grade }).
		Execute()

	fmt.Printf("Active students sorted by grade: %v\n", result)
}
