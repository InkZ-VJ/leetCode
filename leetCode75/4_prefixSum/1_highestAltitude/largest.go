package highestaltitude

func largestAltitude(gain []int) int {
	heightest := 0
	current := 0
	for _, num := range gain {
		current += num
		if current > heightest {
			heightest = current
		}
	}
	return heightest
}
