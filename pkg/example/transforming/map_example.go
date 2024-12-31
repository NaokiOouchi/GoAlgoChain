package transforming

import (
	"fmt"
	"strings"

	"github.com/NaokiOouchi/GoAlgoChain/pkg/algo"
)

func MapBasicExample() {
	numbers := []int{1, 2, 3, 4, 5}
	result, _ := algo.NewPipelineWithData(numbers).
		Map(func(n int) int { return n * n }).
		Execute()

	fmt.Printf("Squared numbers: %v\n", result)
}

type User struct {
	ID       int
	Name     string
	Email    string
	IsActive bool
}

func MapStructExample() {
	users := []User{
		{ID: 1, Name: "alice smith", Email: "alice@example.com", IsActive: true},
		{ID: 2, Name: "bob jones", Email: "bob@example.com", IsActive: true},
	}

	// Capitalize names
	result, _ := algo.NewPipelineWithData(users).
		Map(func(u User) User {
			u.Name = strings.ToUpper(u.Name)
			return u
		}).
		Execute()

	fmt.Printf("Users with capitalized names: %v\n", result)
}

func MapCombinedExample() {
	users := []User{
		{ID: 1, Name: "alice", Email: "alice@example.com", IsActive: true},
		{ID: 2, Name: "bob", Email: "bob@example.com", IsActive: false},
		{ID: 3, Name: "charlie", Email: "charlie@example.com", IsActive: true},
	}

	result, _ := algo.NewPipelineWithData(users).
		Filter(func(u User) bool { return u.IsActive }).
		Map(func(u User) User {
			u.Name = strings.ToUpper(u.Name)
			return u
		}).
		Execute()

	fmt.Printf("Active users with capitalized names: %v\n", result)
}
