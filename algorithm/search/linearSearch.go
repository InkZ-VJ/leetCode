package search

func LinearSearch(arr []int, key int) bool {
	for _, num := range arr {
		if num == key {
			return true
		}
	}
	return false
}
