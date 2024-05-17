package movezeros

func moveZeroes(nums []int) []int {
	num_index := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[num_index] = nums[i]
			num_index++
		}
	}
	for i := num_index; i < len(nums); i++ {
		nums[i] = 0
	}
	return nums
}
