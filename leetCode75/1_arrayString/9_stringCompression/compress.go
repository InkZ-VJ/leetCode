package stringcompression

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	fmt.Println("Input: ", string(chars))
	count, i := 0, 1
	fmt.Printf("index-%d add char %s \n", 0, string(chars[0]))
	top := chars[0]
	for _, b := range chars {
		if b != top {
			fmt.Println("found change")
			if count > 1 {
				strN := strconv.Itoa(count)
				for _, digit := range strN {
					fmt.Printf("index-%d add # %s = %d\n", i, string(top), digit)
					chars[i] = byte(digit)
					i++
				}
			}
			fmt.Printf("index-%d add char %s \n", i, string(b))
			chars[i] = b
			i++
			top = b
			count = 0
		}
		count++
	}
	if count > 1 {
		strN := strconv.Itoa(count)
		for _, digit := range strN {
			fmt.Printf("index-%d add # %s = %d\n", i, string(top), digit)
			chars[i] = byte(digit)
			i++
		}
	}
	chars = chars[:i]
	fmt.Println("final chars: ", string(chars))
	return len(chars)
}
