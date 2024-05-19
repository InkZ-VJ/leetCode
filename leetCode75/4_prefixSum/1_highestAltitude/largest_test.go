package highestaltitude

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargest(t *testing.T) {
	testCases := [][]int{{-5, 1, 5, 0, -7}, {-4, -3, -2, -1, 4, 3, 2}}
	exp := []int{1, 0}
	for i, test := range testCases {
		ans := largestAltitude(test)
		assert.Equal(t, exp[i], ans)
	}
}
