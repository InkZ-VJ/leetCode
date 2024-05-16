package reversevowels

func reverseVowels(s string) string {
	vowels := map[string]bool{"a": true, "e": true, "i": true, "o": true, "u": true}
	i, j := 0, len(s)-1
	input := []rune(s)
	for {

		// is I vowels??
		i_s := string(input[i])
		_, i_ok := vowels[i_s]
		if !i_ok {
			i++
		}

		// is J vowels??
		j_s := string(input[j])
		_, j_ok := vowels[j_s]
		if !j_ok {
			j--
		}

		if i_ok && j_ok {
			temp := input[i]
			input[i] = input[j]
			input[j] = temp
			i++
			j--
		}

		if j <= i {
			break
		}
	}
	return string(input)
}
