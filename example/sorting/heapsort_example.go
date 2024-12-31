package sorting

import (
	"fmt"

	"github.com/NaokiOouchi/GoAlgoChain/pkg/algo"
)

func HeapSortBasicExample() {
	numbers := []int{12, 11, 13, 5, 6, 7}
	result, _ := algo.NewPipelineWithData(numbers).
		HeapSort(func(a, b int) bool { return a < b }).
		Execute()

	fmt.Printf("Sorted numbers: %v\n", result)
}

type Task struct {
	ID       int
	Priority int
	Name     string
}

func HeapSortStructExample() {
	tasks := []Task{
		{ID: 1, Priority: 3, Name: "Debug issue"},
		{ID: 2, Priority: 1, Name: "Write docs"},
		{ID: 3, Priority: 5, Name: "Fix critical bug"},
	}

	// Sort by priority descending
	result, _ := algo.NewPipelineWithData(tasks).
		HeapSort(func(a, b Task) bool { return a.Priority > b.Priority }).
		Execute()

	fmt.Printf("Tasks sorted by priority: %v\n", result)
}

func HeapSortCombinedExample() {
	tasks := []Task{
		{ID: 1, Priority: 3, Name: "Debug issue"},
		{ID: 2, Priority: 1, Name: "Write docs"},
		{ID: 3, Priority: 5, Name: "Fix critical bug"},
		{ID: 4, Priority: 2, Name: "Update readme"},
	}

	result, _ := algo.NewPipelineWithData(tasks).
		Filter(func(t Task) bool { return t.Priority > 2 }).
		HeapSort(func(a, b Task) bool { return a.Priority > b.Priority }).
		Execute()

	fmt.Printf("High priority tasks sorted: %v\n", result)
}
