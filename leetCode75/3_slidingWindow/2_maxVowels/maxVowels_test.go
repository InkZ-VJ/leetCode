package maxvowels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxVowels(t *testing.T) {
	testCases := []string{"abciiidef", "aeiou", "leetcode"}
	k := []int{3, 2, 3}
	exp := []int{3, 2, 2}
	for i, test := range testCases {
		output := maxVowels(test, k[i])
		assert.Equal(t, exp[i], output)
	}
}
