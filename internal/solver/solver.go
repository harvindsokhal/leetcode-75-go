package solver

import "github.com/harvindsokhal/leetcode-75-go/internal/workspace"

func Solve(target string) error {
	return workspace.Open(target, workspace.Options{
		StartTimer: true,
	})
}
