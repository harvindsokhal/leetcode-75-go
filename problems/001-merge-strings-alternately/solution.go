package solution

func mergeAlternately(word1 string, word2 string) string {
	result := make([]byte, 0, len(word1)+len(word2))

	i := 0

	for i < len(word1) && i < len(word2) {
		result = append(result, word1[i])
		result = append(result, word2[i])
		i++
	}

	result = append(result, word1[i:]...)
	result = append(result, word2[i:]...)

	return string(result)
}
