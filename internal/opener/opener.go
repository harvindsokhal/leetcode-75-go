package opener

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/harvindsokhal/leetcode-75-go/internal/problems"
	"github.com/harvindsokhal/leetcode-75-go/internal/tracker"
)

func Open(slug string) error {
	resolvedSlug, err := resolveSlug(slug)
	if err != nil {
		return err
	}

	problem, ok := problems.FindBySlug(resolvedSlug)
	if !ok {
		return fmt.Errorf("problem %q not found", resolvedSlug)
	}
	folder := problems.FolderName(problem)

	solutionPath := fmt.Sprintf(
		"problems/%s/solution.go",
		folder,
	)

	link := fmt.Sprintf(
		"https://leetcode.com/problems/%s/",
		problem.Slug,
	)

	fmt.Println("Opening:", solutionPath)
	fmt.Println("Link:", link)

	if err := openBrowser(link); err != nil {
		return err
	}

	return openInNeovim(solutionPath)
}

func resolveSlug(slug string) (string, error) {
	switch slug {
	case "next":
		return tracker.NextSlug()
	case "current":
		return tracker.CurrentSlug()
	default:
		return slug, nil
	}
}

func openInNeovim(path string) error {
	cmd := exec.Command("nvim", path)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
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
