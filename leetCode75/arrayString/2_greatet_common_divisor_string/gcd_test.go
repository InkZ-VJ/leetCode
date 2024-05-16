package greatetcommondivisorstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	testCases := [3][2]string{{"ABCABC", "ABC"}, {"ABABAB", "ABAB"}, {"LEET", "CODE"}}
	expectedAns := [3]string{"ABC", "AB", ""}
	for i, test := range testCases {
		output := gcdOfStrings(test[0], test[1])
		assert.Equal(t, output, expectedAns[i])
	}

}
