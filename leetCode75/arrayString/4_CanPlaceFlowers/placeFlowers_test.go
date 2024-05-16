package canplaceflowers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaceFlowers(t *testing.T) {
	testCases_1 := [2][]int{{1, 0, 0, 0, 1}, {1, 0, 0, 0, 1}}
	testCases_2 := [2]int{1, 2}
	expectedAns := [2]bool{true, false}

	for i, test := range testCases_1 {
		ans := canPlaceFlowers(test, testCases_2[i])
		assert.Equal(t, ans, expectedAns[i])
	}
}
