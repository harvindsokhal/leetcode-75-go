package main

import (
	"fmt"
	"os"

	"github.com/harvindsokhal/leetcode-75-go/internal/generator"
	githelper "github.com/harvindsokhal/leetcode-75-go/internal/git"
	"github.com/harvindsokhal/leetcode-75-go/internal/problems"
	"github.com/harvindsokhal/leetcode-75-go/internal/runner"
	"github.com/harvindsokhal/leetcode-75-go/internal/tracker"
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

	case "start":
		if len(os.Args) < 3 {
			fmt.Println("Usage: lc75 start <problem-slug>")
			os.Exit(1)
		}

		slug := os.Args[2]

		if err := tracker.Start(slug); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("Started:", slug)

	case "finish":
		if len(os.Args) < 3 {
			fmt.Println("Usage: lc75 finish <problem-slug>")
			os.Exit(1)
		}

		slug := os.Args[2]

		if err := tracker.Finish(slug); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("Finished:", slug)

	case "status":
		if err := tracker.Status(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

	case "in-progress":

		if len(os.Args) < 3 {
			fmt.Println("Usage: lc75 start <problem-slug>")
			os.Exit(1)
		}

		slug := os.Args[2]

		if err := tracker.In_progress(slug); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

	case "test":
		if len(os.Args) < 3 {
			fmt.Println("Usage: lc75 test <problem-slug>")
		}

		slug := os.Args[2]

		if err := runner.Test(slug); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

	case "commit":
		if len(os.Args) < 3 {
			fmt.Println("Usage: lc75 commit <problem-slug>")
			os.Exit(1)
		}

		slug := os.Args[2]

		if err := githelper.CommitProblem(slug); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("Commited:", slug)

	case "push":
		if err := githelper.Push(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("Pushed changes.")

	case "stats":
		if err := tracker.Stats(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

	case "help":
		fmt.Println("  list                  List all problems")
		printHelp()

	case "list":
		printProblems()

	case "current":
		if err := tracker.Current(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

	case "next":
		if err := tracker.Next(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println()
		printHelp()
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

func printHelp() {
	fmt.Println("LeetCode 75 CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  lc75 <command> [args]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  init                  Generate folders/files for all problems")
	fmt.Println("  start <slug>           Start timer for a problem")
	fmt.Println("  finish <slug>          Finish timer for a problem")
	fmt.Println("  current               Show current in-progress problem")
	fmt.Println("  next                  Show next unsolved problem")
	fmt.Println("  status                 Show current progress")
	fmt.Println("  stats                  Show detailed stats")
	fmt.Println("  test <slug>            Run Go tests for a problem")
	fmt.Println("  commit <slug>          Commit a completed problem")
	fmt.Println("  push                   Push commits to remote")
	fmt.Println("  help                   Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  lc75 start merge-strings-alternately")
	fmt.Println("  lc75 test merge-strings-alternately")
	fmt.Println("  lc75 finish merge-strings-alternately")
	fmt.Println("  lc75 commit merge-strings-alternately")
}
