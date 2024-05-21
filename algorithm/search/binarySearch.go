package search

func BinarySearch(arr []int, key int) bool {
	if len(arr) == 1 {
		return arr[0] == key
	}
	middle := len(arr) / 2
	if arr[middle] == key {
		return true
	} else if arr[middle] < key {
		return BinarySearch(arr[middle+1:], key) // greb right side
	} else {
		return BinarySearch(arr[:middle], key) // greb left side
	}
}
