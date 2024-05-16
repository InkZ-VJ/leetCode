package mergestringaltenately

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	testCases := [3][2]string{{"abc", "pqr"}, {"ab", "pqrs"}, {"abcd", "pq"}}
	expectedAns := [3]string{"apbqcr", "apbqrs", "apbqcd"}

	for i, test := range testCases {
		output := mergeAlternately(test[0], test[1])
		assert.Equal(t, output, expectedAns[i])
	}
}
