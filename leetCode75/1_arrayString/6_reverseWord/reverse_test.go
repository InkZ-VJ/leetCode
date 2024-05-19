package reverseword

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	testCases := []string{"the sky is blue", "  hello world  ", "a good   example"}
	expectedAns := []string{"blue is sky the", "world hello", "example good a"}

	for i, test := range testCases {
		ans := reverseWords(test)
		assert.Equal(t, expectedAns[i], ans)
	}
}
