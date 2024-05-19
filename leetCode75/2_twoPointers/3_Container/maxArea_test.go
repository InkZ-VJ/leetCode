package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {
	testCases := [][]int{{1, 8, 6, 2, 5, 4, 8, 3, 7}, {1, 1}}
	exp := []int{49, 1}
	for i, test := range testCases {
		output := maxArea(test)
		assert.Equal(t, exp[i], output)
	}
}
