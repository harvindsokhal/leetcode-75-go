package problems

type Problem struct {
	Number     int
	Title      string
	Slug       string
	Difficulty string
	Category   string
}

var LeetCode75 = []Problem{
	{1, "Merge Strings Alternately", "merge-strings-alternately", "Easy", "Array / String"},
	{2, "Greatest Common Divisor of Strings", "greatest-common-divisor-of-strings", "Easy", "Array / String"},
	{3, "Kids With the Greatest Number of Candies", "kids-with-the-greatest-number-of-candies", "Easy", "Array / String"},
	{4, "Can Place Flowers", "can-place-flowers", "Easy", "Array / String"},
	{5, "Reverse Vowels of a String", "reverse-vowels-of-a-string", "Easy", "Array / String"},
	{6, "Reverse Words in a String", "reverse-words-in-a-string", "Medium", "Array / String"},
	{7, "Product of Array Except Self", "product-of-array-except-self", "Medium", "Array / String"},
	{8, "Increasing Triplet Subsequence", "increasing-triplet-subsequence", "Medium", "Array / String"},
	{9, "String Compression", "string-compression", "Medium", "Array / String"},

	{10, "Move Zeroes", "move-zeroes", "Easy", "Two Pointers"},
	{11, "Is Subsequence", "is-subsequence", "Easy", "Two Pointers"},
	{12, "Container With Most Water", "container-with-most-water", "Medium", "Two Pointers"},
	{13, "Max Number of K-Sum Pairs", "max-number-of-k-sum-pairs", "Medium", "Two Pointers"},

	{14, "Maximum Average Subarray I", "maximum-average-subarray-i", "Easy", "Sliding Window"},
	{15, "Maximum Number of Vowels in a Substring of Given Length", "maximum-number-of-vowels-in-a-substring-of-given-length", "Medium", "Sliding Window"},
	{16, "Max Consecutive Ones III", "max-consecutive-ones-iii", "Medium", "Sliding Window"},
	{17, "Longest Subarray of 1's After Deleting One Element", "longest-subarray-of-1s-after-deleting-one-element", "Medium", "Sliding Window"},

	{18, "Find the Highest Altitude", "find-the-highest-altitude", "Easy", "Prefix Sum"},
	{19, "Find Pivot Index", "find-pivot-index", "Easy", "Prefix Sum"},

	{20, "Find the Difference of Two Arrays", "find-the-difference-of-two-arrays", "Easy", "Hash Map / Set"},
	{21, "Unique Number of Occurrences", "unique-number-of-occurrences", "Easy", "Hash Map / Set"},
	{22, "Determine if Two Strings Are Close", "determine-if-two-strings-are-close", "Medium", "Hash Map / Set"},
	{23, "Equal Row and Column Pairs", "equal-row-and-column-pairs", "Medium", "Hash Map / Set"},

	{24, "Removing Stars From a String", "removing-stars-from-a-string", "Medium", "Stack"},
	{25, "Asteroid Collision", "asteroid-collision", "Medium", "Stack"},
	{26, "Decode String", "decode-string", "Medium", "Stack"},

	{27, "Number of Recent Calls", "number-of-recent-calls", "Easy", "Queue"},
	{28, "Dota2 Senate", "dota2-senate", "Medium", "Queue"},

	{29, "Delete the Middle Node of a Linked List", "delete-the-middle-node-of-a-linked-list", "Medium", "Linked List"},
	{30, "Odd Even Linked List", "odd-even-linked-list", "Medium", "Linked List"},
	{31, "Reverse Linked List", "reverse-linked-list", "Easy", "Linked List"},
	{32, "Maximum Twin Sum of a Linked List", "maximum-twin-sum-of-a-linked-list", "Medium", "Linked List"},

	{33, "Maximum Depth of Binary Tree", "maximum-depth-of-binary-tree", "Easy", "Binary Tree - DFS"},
	{34, "Leaf-Similar Trees", "leaf-similar-trees", "Easy", "Binary Tree - DFS"},
	{35, "Count Good Nodes in Binary Tree", "count-good-nodes-in-binary-tree", "Medium", "Binary Tree - DFS"},
	{36, "Path Sum III", "path-sum-iii", "Medium", "Binary Tree - DFS"},
	{37, "Longest ZigZag Path in a Binary Tree", "longest-zigzag-path-in-a-binary-tree", "Medium", "Binary Tree - DFS"},
	{38, "Lowest Common Ancestor of a Binary Tree", "lowest-common-ancestor-of-a-binary-tree", "Medium", "Binary Tree - DFS"},

	{39, "Binary Tree Right Side View", "binary-tree-right-side-view", "Medium", "Binary Tree - BFS"},
	{40, "Maximum Level Sum of a Binary Tree", "maximum-level-sum-of-a-binary-tree", "Medium", "Binary Tree - BFS"},

	{41, "Search in a Binary Search Tree", "search-in-a-binary-search-tree", "Easy", "Binary Search Tree"},
	{42, "Delete Node in a BST", "delete-node-in-a-bst", "Medium", "Binary Search Tree"},

	{43, "Keys and Rooms", "keys-and-rooms", "Medium", "Graphs - DFS"},
	{44, "Number of Provinces", "number-of-provinces", "Medium", "Graphs - DFS"},
	{45, "Reorder Routes to Make All Paths Lead to the City Zero", "reorder-routes-to-make-all-paths-lead-to-the-city-zero", "Medium", "Graphs - DFS"},
	{46, "Evaluate Division", "evaluate-division", "Medium", "Graphs - DFS"},

	{47, "Nearest Exit from Entrance in Maze", "nearest-exit-from-entrance-in-maze", "Medium", "Graphs - BFS"},
	{48, "Rotting Oranges", "rotting-oranges", "Medium", "Graphs - BFS"},

	{49, "Kth Largest Element in an Array", "kth-largest-element-in-an-array", "Medium", "Heap / Priority Queue"},
	{50, "Smallest Number in Infinite Set", "smallest-number-in-infinite-set", "Medium", "Heap / Priority Queue"},
	{51, "Maximum Subsequence Score", "maximum-subsequence-score", "Medium", "Heap / Priority Queue"},
	{52, "Total Cost to Hire K Workers", "total-cost-to-hire-k-workers", "Medium", "Heap / Priority Queue"},

	{53, "Guess Number Higher or Lower", "guess-number-higher-or-lower", "Easy", "Binary Search"},
	{54, "Successful Pairs of Spells and Potions", "successful-pairs-of-spells-and-potions", "Medium", "Binary Search"},
	{55, "Find Peak Element", "find-peak-element", "Medium", "Binary Search"},
	{56, "Koko Eating Bananas", "koko-eating-bananas", "Medium", "Binary Search"},

	{57, "Letter Combinations of a Phone Number", "letter-combinations-of-a-phone-number", "Medium", "Backtracking"},
	{58, "Combination Sum III", "combination-sum-iii", "Medium", "Backtracking"},

	{59, "N-th Tribonacci Number", "n-th-tribonacci-number", "Easy", "DP - 1D"},
	{60, "Min Cost Climbing Stairs", "min-cost-climbing-stairs", "Easy", "DP - 1D"},
	{61, "House Robber", "house-robber", "Medium", "DP - 1D"},
	{62, "Domino and Tromino Tiling", "domino-and-tromino-tiling", "Medium", "DP - 1D"},

	{63, "Unique Paths", "unique-paths", "Medium", "DP - Multidimensional"},
	{64, "Longest Common Subsequence", "longest-common-subsequence", "Medium", "DP - Multidimensional"},
	{65, "Best Time to Buy and Sell Stock with Transaction Fee", "best-time-to-buy-and-sell-stock-with-transaction-fee", "Medium", "DP - Multidimensional"},
	{66, "Edit Distance", "edit-distance", "Medium", "DP - Multidimensional"},

	{67, "Counting Bits", "counting-bits", "Easy", "Bit Manipulation"},
	{68, "Single Number", "single-number", "Easy", "Bit Manipulation"},
	{69, "Minimum Flips to Make a OR b Equal to c", "minimum-flips-to-make-a-or-b-equal-to-c", "Medium", "Bit Manipulation"},

	{70, "Implement Trie Prefix Tree", "implement-trie-prefix-tree", "Medium", "Trie"},
	{71, "Search Suggestions System", "search-suggestions-system", "Medium", "Trie"},

	{72, "Non-overlapping Intervals", "non-overlapping-intervals", "Medium", "Intervals"},
	{73, "Minimum Number of Arrows to Burst Balloons", "minimum-number-of-arrows-to-burst-balloons", "Medium", "Intervals"},

	{74, "Daily Temperatures", "daily-temperatures", "Medium", "Monotonic Stack"},
	{75, "Online Stock Span", "online-stock-span", "Medium", "Monotonic Stack"},
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
