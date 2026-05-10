package problems

import "fmt"

func FolderName(problem Problem) string {
	return fmt.Sprintf("%03d-%s", problem.Number, problem.Slug)
}
