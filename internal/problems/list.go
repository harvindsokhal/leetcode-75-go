package problems

type Problem struct {
	Number     int
	Title      string
	Slug       string
	Difficulty string
	Category   string
}

var LeetCode75 = []Problem{
	{
		Number:     1,
		Title:      "Merge Strings Alternately",
		Slug:       "merge-strings-alternately",
		Difficulty: "Easy",
		Category:   "Array / String",
	},
	{
		Number:     2,
		Title:      "Greatest Common Divisor of Strings",
		Slug:       "greatest-common-divisor-of-strings",
		Difficulty: "Easy",
		Category:   "Array / String",
	},
	{
		Number:     3,
		Title:      "Kids With the Greatest Number of Candies",
		Slug:       "kids-with-the-greatest-number-of-candies",
		Difficulty: "Easy",
		Category:   "Array / String",
	},
}

func All() []Problem {
	return LeetCode75
}

func FindBySlug(slug string) (Problem, bool) {
	for _, problem := range LeetCode75 {
		if problem.Slug == slug {
			return problem, true
		}
	}

	return Problem{}, false
}
