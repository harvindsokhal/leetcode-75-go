package main

import (
	"fmt"

	"github.com/harvindsokhal/leetcode-75-go/internal/problems"
)

func main() {
	fmt.Println("Leetcode 75 Problems")
	fmt.Println("--------------------")

	for _, problem := range problems.All() {
		fmt.Printf("%03d. %s [%s] - %s\n",
			problem.Number,
			problem.Title,
			problem.Difficulty,
			problem.Category,
		)
	}
}
