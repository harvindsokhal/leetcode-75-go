package workspace

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/harvindsokhal/leetcode-75-go/internal/problems"
	"github.com/harvindsokhal/leetcode-75-go/internal/tracker"
)

type Options struct {
	StartTimer bool
}

func Open(target string, opts Options) error {
	slug, err := resolveSlug(target)
	if err != nil {
		return err
	}

	problem, ok := problems.FindBySlug(slug)
	if !ok {
		return fmt.Errorf("problem %q not found", slug)
	}

	if opts.StartTimer {
		if err := tracker.Start(slug); err != nil {
			fmt.Println("Timer note:", err)
		}
	}

	folder := problems.FolderName(problem)

	solutionPath := fmt.Sprintf("problems/%s/solution.go", folder)
	readmePath := fmt.Sprintf("problems/%s/README.md", folder)
	testPath := fmt.Sprintf("problems/%s/solution_test.go", folder)
	link := fmt.Sprintf("https://leetcode.com/problems/%s/", problem.Slug)

	fmt.Printf("Problem: %03d. %s\n", problem.Number, problem.Title)
	fmt.Println("Link:", link)
	fmt.Println("Workspace:", folder)

	if err := openBrowser(link); err != nil {
		return err
	}

	return openTmuxWorkspace(problem, solutionPath, readmePath, testPath)
}

func resolveSlug(target string) (string, error) {
	switch target {
	case "next":
		return tracker.NextSlug()
	case "current":
		return tracker.CurrentSlug()
	default:
		return target, nil
	}
}

func openBrowser(target string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", target)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", target)
	default:
		cmd = exec.Command("wslview", target)
	}

	return cmd.Start()
}

func openTmuxWorkspace(problem problems.Problem, solutionPath, readmePath, testPath string) error {
	sessionName := fmt.Sprintf("lc75-%03d", problem.Number)

	nvimCommand := fmt.Sprintf(
		"nvim %s -c 'vsplit %s' -c 'split %s' -c 'wincmd h'",
		solutionPath,
		readmePath,
		testPath,
	)

	cmd := exec.Command(
		"tmux",
		"new-session",
		"-A",
		"-s",
		sessionName,
		nvimCommand,
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
