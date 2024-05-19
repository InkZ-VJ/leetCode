package findpivot

func pivotIndex(nums []int) int {
	ans := -1
	leftSum, rightSum := 0, 0

	for i := 1; i < len(nums); i++ {
		rightSum += nums[i]
	}
	if leftSum == rightSum {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		leftSum += nums[i-1]
		rightSum -= nums[i]
		if leftSum == rightSum {
			return i
		}
	}
	return ans
}
