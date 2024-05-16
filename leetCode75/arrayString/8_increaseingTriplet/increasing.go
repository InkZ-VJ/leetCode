package increaseingtriplet

func increasingTriplet(nums []int) bool {
	for i := 0; i < len(nums)-2; i++ {
		// fmt.Println(i)
		if nums[i] < nums[i+1] && nums[i+1] < nums[i+2] {
			return true
		}
	}
	return false
}
