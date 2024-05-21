package sorting

func merge(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	middle := len(arr) / 2
	left := arr[:middle]
	right := arr[middle:]

	arr = mergesort(merge(left), merge(right))
	return arr
}

func mergesort(left []int, right []int) []int {
	var response []int
	lIndex, rIndex := 0, 0
	for lIndex < len(left) || rIndex < len(right) {
		if lIndex < len(left) && (rIndex == len(right) || left[lIndex] < right[rIndex]) {
			response = append(response, left[lIndex])
			lIndex++
		} else {
			response = append(response, right[rIndex])
			rIndex++
		}
	}
	return response
}
