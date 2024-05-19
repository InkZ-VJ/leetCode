package maxaverage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	testCases := [][]int{{1, 12, -5, -6, 50, 3}, {5}}
	k := []int{4, 1}
	exp := []float64{12.750, 5.0}
	for i, test := range testCases {
		output := findMaxAverage(test, k[i])
		assert.Equal(t, exp[i], output)
	}
}
