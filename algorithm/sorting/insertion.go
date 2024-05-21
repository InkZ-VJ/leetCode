package sorting

func insertion(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		if arr[0] > arr[i] {
			// arr[i] go to first index
			temp := arr[i]
			copy(arr[1:i+1], arr[0:i])
			arr[0] = temp
		} else {
			// swap until met condition
			for j := 1; j < i; j++ {
				if arr[j-1] <= arr[i] && arr[j] > arr[i] {
					arr[j], arr[i] = arr[i], arr[j]
				}
			}
		}
	}
	return arr
}
