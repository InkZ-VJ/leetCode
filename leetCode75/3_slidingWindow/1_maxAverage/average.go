package maxaverage

func findMaxAverage(nums []int, k int) float64 {
	maxSum := int(-1 << 63)
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	if sum > maxSum {
		maxSum = sum
	}
	if len(nums) == k {
		return float64(sum) / float64(k)
	}

	tail := k
	for head := 1; head < len(nums)-k+1; head++ {
		sum += (nums[tail] - nums[head-1])
		tail++
		if maxSum < sum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}
