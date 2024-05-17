package container

func maxArea(height []int) int {
	ans, i, j := 0, 0, len(height)-1
	for {
		if i >= j {
			break
		}

		tmp := (j - i)
		if height[i] < height[j] {
			tmp *= height[i]
		} else {
			tmp *= height[j]
		}
		// fmt.Println("current tmp: ", tmp)
		if tmp > ans {
			ans = tmp
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
		// fmt.Printf("break at haed: %d, tail: %d\n", i, j)
	}
	return ans
}
