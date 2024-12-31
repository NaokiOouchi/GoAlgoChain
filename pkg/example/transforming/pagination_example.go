package transforming

import (
	"fmt"

	"github.com/NaokiOouchi/GoAlgoChain/pkg/algo"
)

func SkipTakeBasicExample() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result, _ := algo.NewPipelineWithData(numbers).
		Skip(3).
		Take(4).
		Execute()

	fmt.Printf("Page of numbers: %v\n", result) // [4, 5, 6, 7]
}

type Article struct {
	ID    int
	Title string
	Views int
}

func PaginationStructExample() {
	articles := []Article{
		{ID: 1, Title: "First Post", Views: 100},
		{ID: 2, Title: "Second Post", Views: 200},
		{ID: 3, Title: "Third Post", Views: 300},
		{ID: 4, Title: "Fourth Post", Views: 400},
		{ID: 5, Title: "Fifth Post", Views: 500},
	}

	// Get second page (2 items per page)
	page2, _ := algo.NewPipelineWithData(articles).
		Skip(2).
		Take(2).
		Execute()

	fmt.Printf("Page 2 of articles: %v\n", page2)
}

func PaginationCombinedExample() {
	articles := []Article{
		{ID: 1, Title: "First Post", Views: 100},
		{ID: 2, Title: "Second Post", Views: 200},
		{ID: 3, Title: "Third Post", Views: 300},
		{ID: 4, Title: "Fourth Post", Views: 400},
		{ID: 5, Title: "Fifth Post", Views: 500},
	}

	result, _ := algo.NewPipelineWithData(articles).
		QuickSort(func(a, b Article) bool { return a.Views > b.Views }).
		Skip(1).
		Take(3).
		Execute()

	fmt.Printf("Top viewed articles (2-4): %v\n", result)
}
