package main

import (
	"fmt"
)

func main() {
	testCases := [][]int{{1, 12, -5, -6, 50, 3}, {5}}
	k := []int{4, 1}
	exp := []float64{12.750, 5.0}
	for i, test := range testCases {
		fmt.Println("Input: ", test)
		fmt.Println("Expected: ", exp[i])
		output := findMaxAverage(test, k[i])
		fmt.Println("Output: ", output)
	}
}

func findMaxAverage(nums []int, k int) float64 {
	head, tail := 0, k-1
	ans := float64(0)
	for tail < len(nums) {
		avg := findAVG(nums[head : tail+1])
		if avg > ans {
			ans = avg
		}
		head++
		tail++
	}
	return ans
}

func findAVG(nums []int) float64 {
	avg := 0
	for _, num := range nums {
		avg += num
	}
	return float64(avg) / float64(len(nums))
}
