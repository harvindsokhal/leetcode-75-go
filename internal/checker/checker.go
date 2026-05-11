package checker

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/harvindsokhal/leetcode-75-go/internal/problems"
	"github.com/harvindsokhal/leetcode-75-go/internal/resolver"
)

func Check(target string) error {
	slug, err := resolver.Resolve(target)
	if err != nil {
		return err
	}

	problem, ok := problems.FindBySlug(slug)
	if !ok {
		return fmt.Errorf("problem %q not found", slug)
	}

	fmt.Printf("Checking %03d - %s\n", problem.Number, problem.Title)

	if err := runGoFmt(problem); err != nil {
		return err
	}

	if err := runGoTest(problem); err != nil {
		return err
	}

	fmt.Println("Check passed.")
	return nil
}

func runGoFmt(problem problems.Problem) error {
	folder := problems.FolderName(problem)
	path := fmt.Sprintf("./problems/%s", folder)

	cmd := exec.Command("go", "fmt", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func runGoTest(problem problems.Problem) error {
	folder := problems.FolderName(problem)
	path := fmt.Sprintf("./problems/%s", folder)

	cmd := exec.Command("go", "test", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
