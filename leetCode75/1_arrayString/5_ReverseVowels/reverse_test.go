package reversevowels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	testCases := [2]string{"hello", "leetcode"}
	expectedAns := [2]string{"holle", "leotcede"}
	for i, test := range testCases {
		ans := reverseVowels(test)
		assert.Equal(t, ans, expectedAns[i])
	}
}
