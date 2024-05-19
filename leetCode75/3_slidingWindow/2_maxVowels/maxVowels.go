package maxvowels

func maxVowels(s string, k int) int {
	maxCount, count := 0, 0
	vowels := map[string]bool{"a": true, "e": true, "i": true, "o": true, "u": true}
	for i := 0; i < k; i++ {
		_, ok := vowels[string(s[i])]
		if ok {
			count++
		}
	}
	if count > maxCount {
		maxCount = count
	}
	// fmt.Println("first window count: ", count)

	tail := k
	for head := 1; head < len(s)-k+1; head++ {
		// fmt.Printf("window %d->%d: %v\n", head, tail, s[head:tail+1])
		_, ok := vowels[string(s[head-1])]
		if ok {
			count--
		}
		_, ok = vowels[string(s[tail])]
		if ok {
			count++
		}
		// fmt.Printf("window count: %d\n", count)
		if count > maxCount {
			maxCount = count
		}
		tail++
	}
	return maxCount
}
