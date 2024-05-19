package reverseword

func reverseWords(s string) string {
	ans := ""
	input := []rune(s)
	head, tail := len(s)-1, len(s)-1
	for {
		if tail < 0 {
			break
		}
		// fmt.Println("current tail: ", tail)
		if string(input[tail]) == " " {
			tail--
			continue
		}
		// fmt.Println("Found tail of word at index: ", tail)
		head = tail - 1
		for {
			if head <= 0 {
				head = 0
				break
			}
			if string(input[head]) != " " {
				head--
			} else {
				break
			}
		}
		// fmt.Println("Found head of word at index: ", head)

		if head == 0 {
			if string(input[head]) != " " {
				ans += string(input[head:tail+1]) + " "
			} else {
				break
			}
			break
		} else {
			ans += string(input[head+1:tail+1]) + " "
		}

		tail = head
	}
	return ans[:len(ans)-1]
}
