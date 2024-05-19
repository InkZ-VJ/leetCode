package stringcompression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompass(t *testing.T) {
	testCases := [][]byte{[]byte("aabbccc"), []byte("a"), []byte("abbbbbbbbbbb")}
	exp := []int{6, 1, 4}
	for i, test := range testCases {
		ans := compress(test)
		assert.Equal(t, exp[i], ans)
	}
}
