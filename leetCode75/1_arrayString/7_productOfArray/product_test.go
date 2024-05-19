package productofarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	testCases := [][]int{{1, 2, 3, 4}, {-1, 1, 0, -3, 3}}
	exp := [][]int{{24, 12, 8, 6}, {0, 0, 9, 0, 0}}
	for i, test := range testCases {
		ans := productExceptSelf(test)
		assert.Equal(t, exp[i], ans)
	}
}
