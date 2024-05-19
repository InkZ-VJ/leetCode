package kidswiththegreatestnumberofcandies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCandies(t *testing.T) {
	testCases := [3][]int{{2, 3, 5, 1, 3}, {4, 2, 1, 1, 2}, {12, 1, 12}}
	testCases_extra := [3]int{3, 1, 10}
	expectedAns := [3][]bool{{true, true, true, false, true}, {true, false, false, false, false}, {true, false, true}}
	for i, test := range testCases {
		ans := kidsWithCandies(test, testCases_extra[i])
		assert.Equal(t, ans, expectedAns[i])
	}
}
