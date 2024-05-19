package findpivot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivot(t *testing.T) {
	testCases := [][]int{{1, 7, 3, 6, 5, 6}, {1, 2, 3}, {2, 1, -1}}
	exp := []int{3, -1, 0}
	for i, test := range testCases {
		ans := pivotIndex(test)
		assert.Equal(t, exp[i], ans)
	}
}
