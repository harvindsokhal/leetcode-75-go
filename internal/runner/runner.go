package runner

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/harvindsokhal/leetcode-75-go/internal/problems"
)

func Test(slug string) error {
	problem, ok := problems.FindBySlug(slug)
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
