package problems

type Problem struct {
	Number       int
	Title        string
	Slug         string
	Difficulty   string
	Category     string
	FunctionName string
}

var LeetCode75 = []Problem{
	{1, "Merge Strings Alternately", "merge-strings-alternately", "Easy", "Array / String", "mergeAlternately"},
	{2, "Greatest Common Divisor of Strings", "greatest-common-divisor-of-strings", "Easy", "Array / String", "gcdOfStrings"},
	{3, "Kids With the Greatest Number of Candies", "kids-with-the-greatest-number-of-candies", "Easy", "Array / String", "kidsWithCandies"},
	{4, "Can Place Flowers", "can-place-flowers", "Easy", "Array / String", "TODO"},
	{5, "Reverse Vowels of a String", "reverse-vowels-of-a-string", "Easy", "Array / String", "TODO"},
	{6, "Reverse Words in a String", "reverse-words-in-a-string", "Medium", "Array / String", "TODO"},
	{7, "Product of Array Except Self", "product-of-array-except-self", "Medium", "Array / String", "TODO"},
	{8, "Increasing Triplet Subsequence", "increasing-triplet-subsequence", "Medium", "Array / String", "TODO"},
	{9, "String Compression", "string-compression", "Medium", "Array / String", "TODO"},
	{10, "Move Zeroes", "move-zeroes", "Easy", "Two Pointers", "TODO"},
	{11, "Is Subsequence", "is-subsequence", "Easy", "Two Pointers", "TODO"},
	{12, "Container With Most Water", "container-with-most-water", "Medium", "Two Pointers", "TODO"},
	{13, "Max Number of K-Sum Pairs", "max-number-of-k-sum-pairs", "Medium", "Two Pointers", "TODO"},
	{14, "Maximum Average Subarray I", "maximum-average-subarray-i", "Easy", "Sliding Window", "TODO"},
	{15, "Maximum Number of Vowels in a Substring of Given Length", "maximum-number-of-vowels-in-a-substring-of-given-length", "Medium", "Sliding Window", "TODO"},
	{16, "Max Consecutive Ones III", "max-consecutive-ones-iii", "Medium", "Sliding Window", "TODO"},
	{17, "Longest Subarray of 1's After Deleting One Element", "longest-subarray-of-1s-after-deleting-one-element", "Medium", "Sliding Window", "TODO"},
	{18, "Find the Highest Altitude", "find-the-highest-altitude", "Easy", "Prefix Sum", "TODO"},
	{19, "Find Pivot Index", "find-pivot-index", "Easy", "Prefix Sum", "TODO"},
	{20, "Find the Difference of Two Arrays", "find-the-difference-of-two-arrays", "Easy", "Hash Map / Set", "TODO"},
	{21, "Unique Number of Occurrences", "unique-number-of-occurrences", "Easy", "Hash Map / Set", "TODO"},
	{22, "Determine if Two Strings Are Close", "determine-if-two-strings-are-close", "Medium", "Hash Map / Set", "TODO"},
	{23, "Equal Row and Column Pairs", "equal-row-and-column-pairs", "Medium", "Hash Map / Set", "TODO"},
	{24, "Removing Stars From a String", "removing-stars-from-a-string", "Medium", "Stack", "TODO"},
	{25, "Asteroid Collision", "asteroid-collision", "Medium", "Stack", "TODO"},
	{26, "Decode String", "decode-string", "Medium", "Stack", "TODO"},
	{27, "Number of Recent Calls", "number-of-recent-calls", "Easy", "Queue", "TODO"},
	{28, "Dota2 Senate", "dota2-senate", "Medium", "Queue", "TODO"},
	{29, "Delete the Middle Node of a Linked List", "delete-the-middle-node-of-a-linked-list", "Medium", "Linked List", "TODO"},
	{30, "Odd Even Linked List", "odd-even-linked-list", "Medium", "Linked List", "TODO"},
	{31, "Reverse Linked List", "reverse-linked-list", "Easy", "Linked List", "TODO"},
	{32, "Maximum Twin Sum of a Linked List", "maximum-twin-sum-of-a-linked-list", "Medium", "Linked List", "TODO"},
	{33, "Maximum Depth of Binary Tree", "maximum-depth-of-binary-tree", "Easy", "Binary Tree - DFS", "TODO"},
	{34, "Leaf-Similar Trees", "leaf-similar-trees", "Easy", "Binary Tree - DFS", "TODO"},
	{35, "Count Good Nodes in Binary Tree", "count-good-nodes-in-binary-tree", "Medium", "Binary Tree - DFS", "TODO"},
	{36, "Path Sum III", "path-sum-iii", "Medium", "Binary Tree - DFS", "TODO"},
	{37, "Longest ZigZag Path in a Binary Tree", "longest-zigzag-path-in-a-binary-tree", "Medium", "Binary Tree - DFS", "TODO"},
	{38, "Lowest Common Ancestor of a Binary Tree", "lowest-common-ancestor-of-a-binary-tree", "Medium", "Binary Tree - DFS", "TODO"},
	{39, "Binary Tree Right Side View", "binary-tree-right-side-view", "Medium", "Binary Tree - BFS", "TODO"},
	{40, "Maximum Level Sum of a Binary Tree", "maximum-level-sum-of-a-binary-tree", "Medium", "Binary Tree - BFS", "TODO"},
	{41, "Search in a Binary Search Tree", "search-in-a-binary-search-tree", "Easy", "Binary Search Tree", "TODO"},
	{42, "Delete Node in a BST", "delete-node-in-a-bst", "Medium", "Binary Search Tree", "TODO"},
	{43, "Keys and Rooms", "keys-and-rooms", "Medium", "Graphs - DFS", "TODO"},
	{44, "Number of Provinces", "number-of-provinces", "Medium", "Graphs - DFS", "TODO"},
	{45, "Reorder Routes to Make All Paths Lead to the City Zero", "reorder-routes-to-make-all-paths-lead-to-the-city-zero", "Medium", "Graphs - DFS", "TODO"},
	{46, "Evaluate Division", "evaluate-division", "Medium", "Graphs - DFS", "TODO"},
	{47, "Nearest Exit from Entrance in Maze", "nearest-exit-from-entrance-in-maze", "Medium", "Graphs - BFS", "TODO"},
	{48, "Rotting Oranges", "rotting-oranges", "Medium", "Graphs - BFS", "TODO"},
	{49, "Kth Largest Element in an Array", "kth-largest-element-in-an-array", "Medium", "Heap / Priority Queue", "TODO"},
	{50, "Smallest Number in Infinite Set", "smallest-number-in-infinite-set", "Medium", "Heap / Priority Queue", "TODO"},
	{51, "Maximum Subsequence Score", "maximum-subsequence-score", "Medium", "Heap / Priority Queue", "TODO"},
	{52, "Total Cost to Hire K Workers", "total-cost-to-hire-k-workers", "Medium", "Heap / Priority Queue", "TODO"},
	{53, "Guess Number Higher or Lower", "guess-number-higher-or-lower", "Easy", "Binary Search", "TODO"},
	{54, "Successful Pairs of Spells and Potions", "successful-pairs-of-spells-and-potions", "Medium", "Binary Search", "TODO"},
	{55, "Find Peak Element", "find-peak-element", "Medium", "Binary Search", "TODO"},
	{56, "Koko Eating Bananas", "koko-eating-bananas", "Medium", "Binary Search", "TODO"},
	{57, "Letter Combinations of a Phone Number", "letter-combinations-of-a-phone-number", "Medium", "Backtracking", "TODO"},
	{59, "N-th Tribonacci Number", "n-th-tribonacci-number", "Easy", "DP - 1D", "TODO"},
	{60, "Min Cost Climbing Stairs", "min-cost-climbing-stairs", "Easy", "DP - 1D", "TODO"},
	{61, "House Robber", "house-robber", "Medium", "DP - 1D", "TODO"},
	{62, "Domino and Tromino Tiling", "domino-and-tromino-tiling", "Medium", "DP - 1D", "TODO"},
	{63, "Unique Paths", "unique-paths", "Medium", "DP - Multidimensional", "TODO"},
	{64, "Longest Common Subsequence", "longest-common-subsequence", "Medium", "DP - Multidimensional", "TODO"},
	{65, "Best Time to Buy and Sell Stock with Transaction Fee", "best-time-to-buy-and-sell-stock-with-transaction-fee", "Medium", "DP - Multidimensional", "TODO"},
	{66, "Edit Distance", "edit-distance", "Medium", "DP - Multidimensional", "TODO"},
	{67, "Counting Bits", "counting-bits", "Easy", "Bit Manipulation", "TODO"},
	{68, "Single Number", "single-number", "Easy", "Bit Manipulation", "TODO"},
	{69, "Minimum Flips to Make a OR b Equal to c", "minimum-flips-to-make-a-or-b-equal-to-c", "Medium", "Bit Manipulation", "TODO"},
	{70, "Implement Trie Prefix Tree", "implement-trie-prefix-tree", "Medium", "Trie", "TODO"},
	{71, "Search Suggestions System", "search-suggestions-system", "Medium", "Trie", "TODO"},
	{72, "Non-overlapping Intervals", "non-overlapping-intervals", "Medium", "Intervals", "TODO"},
	{73, "Minimum Number of Arrows to Burst Balloons", "minimum-number-of-arrows-to-burst-balloons", "Medium", "Intervals", "TODO"},
	{74, "Daily Temperatures", "daily-temperatures", "Medium", "Monotonic Stack", "TODO"},
	{75, "Online Stock Span", "online-stock-span", "Medium", "Monotonic Stack", "TODO"},
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
