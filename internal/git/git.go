package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/harvindsokhal/leetcode-75-go/internal/problems"
	"github.com/harvindsokhal/leetcode-75-go/internal/resolver"
	"github.com/harvindsokhal/leetcode-75-go/internal/tracker"
)

func CommitProblem(slug string) error {
	resolvedSlug, err := resolver.Resolve(slug)
	if err != nil {
		return err
	}

	problem, ok := problems.FindBySlug(resolvedSlug)
	if !ok {
		return fmt.Errorf("problem %q not found", slug)
	}

	completed, err := tracker.IsCompleted(problem.Slug)
	if err != nil {
		return err
	}

	if !completed {
		return fmt.Errorf("problem %q is not completed yet. Run lc75 finish %s first", problem.Slug, problem.Slug)
	}

	folder := problems.FolderName(problem)
	problemPath := filepath.Join("problems", folder)

	message := fmt.Sprintf("Solve %03d - %s", problem.Number, problem.Title)

	if err := runGit("add", problemPath, "progress.json"); err != nil {
		return err
	}

	if err := runGit("commit", "-m", message); err != nil {
		return err
	}

	return nil
}

func runGit(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Push() error {
	return runGit("push")
}
