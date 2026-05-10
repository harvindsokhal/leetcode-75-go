package main

import (
	"fmt"
	"os"

	"github.com/harvindsokhal/leetcode-75-go/internal/generator"
	"github.com/harvindsokhal/leetcode-75-go/internal/problems"
)

func main() {
	if len(os.Args) < 2 {
		printProblems()
		return
	}

	command := os.Args[1]

	switch command {
	case "init":
		if err := generator.GenerateAll(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("Generated problem folders.")
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}

func printProblems() {
	fmt.Println("LeetCode 75 Problems")
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
