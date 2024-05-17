package maxksum

func maxOperations(nums []int, k int) int {
	n, ans := len(nums), 0

	//base case
	if n <= 1 {
		return ans
	}

	//num : count
	hashmap := make(map[int]int)

	for i := 0; i < n; i++ {

		//skip the values greater or equal to k
		if nums[i] >= k {
			continue
		}

		//check if counterpart of nums[i] is present
		counterpart := k - nums[i]

		if val, exists := hashmap[counterpart]; exists {
			if val != 0 {
				ans++
			}

			//taking care of duplicates
			//example. [1,1,1,2]
			if val-1 == 0 {
				delete(hashmap, counterpart)
			} else {
				hashmap[counterpart]--
			}

		} else {
			hashmap[nums[i]]++
		}

	}
	return ans
}
