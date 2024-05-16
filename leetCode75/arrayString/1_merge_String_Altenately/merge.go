package mergestringaltenately

func mergeAlternately(word1 string, word2 string) string {
	ans := ""
	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		ans += string(word1[i]) + string(word2[j])
		i++
		j++
	}

	for i < len(word1) {
		ans += string(word1[i])
		i++
	}

	for j < len(word2) {
		ans += string(word2[j])
		j++
	}

	return ans
}
