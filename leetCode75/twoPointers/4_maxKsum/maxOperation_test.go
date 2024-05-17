package maxksum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	testCases := [][]int{{1, 2, 3, 4}, {3, 1, 3, 4, 3}}
	k := []int{5, 6}
	exp := []int{2, 1}
	for i, test := range testCases {
		output := maxOperations(test, k[i])
		assert.Equal(t, exp[i], output)
	}
}
