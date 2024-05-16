package kidswiththegreatestnumberofcandies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	// find most candies
	biggest := 0
	for _, amount := range candies {
		if amount > biggest {
			biggest = amount
		}
	}

	ans := make([]bool, len(candies))

	for i, amount := range candies {
		if amount+extraCandies >= biggest {
			ans[i] = true
		} else {
			ans[i] = false
		}
	}
	return ans
}
