package movezeros

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	testCases := [][]int{{0, 1, 0, 3, 12}, {0}}
	exp := [][]int{{1, 3, 12, 0, 0}, {0}}
	for i, test := range testCases {
		after := moveZeroes(test)
		assert.Equal(t, exp[i], after)
	}
}
