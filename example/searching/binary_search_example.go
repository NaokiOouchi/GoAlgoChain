package searching

import (
	"fmt"

	"github.com/NaokiOouchi/GoAlgoChain/pkg/algo"
)

func BinarySearchBasicExample() {
	numbers := []int{3, 7, 12, 19, 23, 45, 78}
	result, _ := algo.NewPipelineWithData(numbers).
		BinarySearch(func(n int) bool { return n == 23 }).
		Execute()

	fmt.Printf("Found number in sorted array: %v\n", result)
}

type User struct {
	ID       int
	Username string
	Role     string
}

func BinarySearchWithSortExample() {
	users := []User{
		{ID: 5, Username: "alice", Role: "admin"},
		{ID: 2, Username: "bob", Role: "user"},
		{ID: 8, Username: "charlie", Role: "user"},
	}

	result, _ := algo.NewPipelineWithData(users).
		QuickSort(func(a, b User) bool { return a.ID < b.ID }).
		BinarySearch(func(u User) bool { return u.ID == 5 }).
		Execute()

	fmt.Printf("Found user: %v\n", result)
}

func BinarySearchCombinedExample() {
	users := []User{
		{ID: 5, Username: "alice", Role: "admin"},
		{ID: 2, Username: "bob", Role: "user"},
		{ID: 8, Username: "charlie", Role: "user"},
		{ID: 3, Username: "dave", Role: "user"},
	}

	result, _ := algo.NewPipelineWithData(users).
		Filter(func(u User) bool { return u.Role == "user" }).
		QuickSort(func(a, b User) bool { return a.ID < b.ID }).
		BinarySearch(func(u User) bool { return u.ID == 3 }).
		Execute()

	fmt.Printf("Found user with role filter: %v\n", result)
}
