package increaseingtriplet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreasing(t *testing.T) {
	testCases := [][]int{{1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}, {2, 1, 5, 0, 4, 6}}
	exp := []bool{true, false, true}
	for i, test := range testCases {
		ans := increasingTriplet(test)
		assert.Equal(t, exp[i], ans)
	}
}
