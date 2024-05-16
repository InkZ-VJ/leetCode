package canplaceflowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	plantable := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] != 0 {
			continue
		}
		// left side case
		if i == 0 && flowerbed[1] == 0 {
			plantable++
			i++
		}
		// right side case
		if i == len(flowerbed) && flowerbed[len(flowerbed)-1] == 0 {
			plantable++
			i++
		}
		if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			plantable++
			i++
		}
	}

	if plantable < n {
		return false
	}
	return true
}
