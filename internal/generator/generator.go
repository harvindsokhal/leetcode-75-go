package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/harvindsokhal/leetcode-75-go/internal/problems"
)

func GenerateAll() error {
	for _, problem := range problems.All() {
		if err := GenerateProblem(problem); err != nil {
			return err
		}
	}

	return nil
}

func GenerateProblem(problem problems.Problem) error {
	folderName := problems.FolderName(problem)
	folderPath := filepath.Join("problems", folderName)

	if err := os.MkdirAll(folderPath, 0o755); err != nil {
		return err
	}

	files := map[string]string{
		"README.md":        readmeTemplate(problem),
		"solution.go":      solutionTemplate(problem),
		"solution_test.go": testTemplate(problem),
		"notes.md":         notesTemplate(),
		"metadata.json":    metadataTemplate(problem),
	}

	for filename, content := range files {
		path := filepath.Join(folderPath, filename)

		if _, err := os.Stat(path); err == nil {
			continue
		}

		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			return err
		}
	}

	return nil
}

func readmeTemplate(problem problems.Problem) string {
	return fmt.Sprintf(`# %03d. %s

## Difficulty
%s

## Category
%s

## Link
https://leetcode.com/problems/%s/

## Function
%s

## Problem Summary
TODO: Summarise the problem in your own words.

## Examples
TODO: Add examples from LeetCode.

## Approach
TODO: Explain your thinking before coding.

## Edge Cases
- TODO

## Complexity
Time:
Space:

## Reflection
What did I learn?
`, problem.Number, problem.Title, problem.Difficulty, problem.Category, problem.Slug, problem.FunctionName)
}

func solutionTemplate(problem problems.Problem) string {
	functionName := problem.FunctionName
	if functionName == "" || functionName == "TODO" {
		functionName = "solution"
	}

	return fmt.Sprintf(`package solution

// %s solves:
// https://leetcode.com/problems/%s/
func %s() {
	// TODO: implement solution
}
`, functionName, problem.Slug, functionName)
}

func testTemplate(problem problems.Problem) string {
	functionName := problem.FunctionName
	if functionName == "" || functionName == "TODO" {
		functionName = "solution"
	}

	return fmt.Sprintf(`package solution

import "testing"

func Test%s(t *testing.T) {
	t.Skip("TODO: add tests for %s")
}
`, exportedTestName(functionName), problem.Title)
}

func exportedTestName(name string) string {
	if name == "" {
		return "Solution"
	}

	return strings.ToUpper(name[:1]) + name[1:]
}

func metadataTemplate(problem problems.Problem) string {
	data, err := json.MarshalIndent(problem, "", "  ")
	if err != nil {
		return "{}"
	}

	return string(data) + "\n"
}

func notesTemplate() string {
	return `# Notes

## First Thoughts

## Pattern Recognition

## Mistakes / Bugs

## Final Explanation

## What I Would Do Differently
`
}
