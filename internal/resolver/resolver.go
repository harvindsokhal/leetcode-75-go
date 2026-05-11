package resolver

import "github.com/harvindsokhal/leetcode-75-go/internal/tracker"

func Resolve(target string) (string, error) {
	switch target {
	case "current":
		return tracker.CurrentSlug()
	case "next":
		return tracker.NextSlug()
	default:
		return target, nil
	}
}
