package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

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
	return fmt.Sprintf(`# %s

## Difficulty
%s

## Category
%s

## Link
https://leetcode.com/problems/%s/

## Problem
TODO: Add problem summary.

## Approach
TODO: Write your thinking here.

## Complexity
Time:
Space:
`, problem.Title, problem.Difficulty, problem.Category, problem.Slug)
}

func solutionTemplate(problem problems.Problem) string {
	return fmt.Sprintf(`package solution

// TODO: Implement solution for %s.
`, problem.Title)
}

func testTemplate(problem problems.Problem) string {
	return fmt.Sprintf(`package solution

import "testing"

func TestSolution(t *testing.T) {
	t.Skip("TODO: add tests for %s")
}
`, problem.Title)
}

func notesTemplate() string {
	return `# Notes

## What I learned

## Mistakes

## Patterns
`
}

func metadataTemplate(problem problems.Problem) string {
	data, err := json.MarshalIndent(problem, "", "  ")
	if err != nil {
		return "{}"
	}

	return string(data) + "\n"
}
