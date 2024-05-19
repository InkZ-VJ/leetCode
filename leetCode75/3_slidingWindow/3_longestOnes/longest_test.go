package longestones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestOnes(t *testing.T) {
	testCases := [][]int{{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, {0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}}
	k := []int{2, 3}
	exp := []int{6, 10}
	for i, test := range testCases {
		output := longestOnes(test, k[i])
		assert.Equal(t, exp[i], output)
	}
}
