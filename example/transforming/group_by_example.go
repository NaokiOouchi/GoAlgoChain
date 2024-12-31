package transforming

import (
	"fmt"

	"github.com/NaokiOouchi/GoAlgoChain/pkg/algo"
)

func GroupByBasicExample() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	result := algo.GroupBy(numbers, func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	})

	fmt.Printf("Numbers grouped by even/odd: %v\n", result)
}

type Student struct {
	ID      int
	Name    string
	Grade   string
	Subject string
}

func GroupByStructExample() {
	students := []Student{
		{ID: 1, Name: "Alice", Grade: "A", Subject: "Math"},
		{ID: 2, Name: "Bob", Grade: "B", Subject: "Math"},
		{ID: 3, Name: "Charlie", Grade: "A", Subject: "Physics"},
	}

	// Group by grade
	byGrade := algo.GroupBy(students, func(s Student) string {
		return s.Grade
	})

	fmt.Printf("Students grouped by grade: %v\n", byGrade)
}

func GroupByCombinedExample() {
	students := []Student{
		{ID: 1, Name: "Alice", Grade: "A", Subject: "Math"},
		{ID: 2, Name: "Bob", Grade: "B", Subject: "Math"},
		{ID: 3, Name: "Charlie", Grade: "A", Subject: "Physics"},
	}

	filtered, _ := algo.NewPipelineWithData(students).
		Filter(func(s Student) bool { return s.Grade == "A" }).
		Execute()

	bySubject := algo.GroupBy(filtered, func(s Student) string {
		return s.Subject
	})

	fmt.Printf("A-grade students grouped by subject: %v\n", bySubject)
}
