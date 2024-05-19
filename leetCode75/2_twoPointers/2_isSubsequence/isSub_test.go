package issubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	testCases := [][]string{{"abc", "ahbgdc"}, {"axc", "ahbgdc"}}
	exp := []bool{true, false}
	for i, test := range testCases {
		output := isSubsequence(test[0], test[1])
		assert.Equal(t, exp[i], output)
	}
}
