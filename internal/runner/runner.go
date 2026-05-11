package runner

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/harvindsokhal/leetcode-75-go/internal/problems"
	"github.com/harvindsokhal/leetcode-75-go/internal/resolver"
)

func Test(slug string) error {
	resolvedSlug, err := resolver.Resolve(slug)
	if err != nil {
		return err
	}

	problem, ok := problems.FindBySlug(resolvedSlug)
	if !ok {
		return fmt.Errorf("problem %q not found", slug)
	}

	folder := problems.FolderName(problem)
	path := fmt.Sprintf("./problems/%s", folder)

	cmd := exec.Command("go", "test", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
