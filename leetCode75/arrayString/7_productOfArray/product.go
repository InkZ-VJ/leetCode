package productofarray

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	for i := range nums {
		ans[i] = 1
		for j := 0; j < i; j++ {
			ans[i] *= nums[j]
		}
		for j := i + 1; j < n; j++ {
			ans[i] *= nums[j]
		}
	}
	return ans
}
