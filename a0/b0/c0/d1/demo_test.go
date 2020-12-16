// test run
// author: baoqiang
// time: 2020/12/15 8:54 下午
package d1

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestDemo(t *testing.T) {
	//testGetContainerWithMostWater(t)
	//testRomanToInteger(t)
	testThreeSum(t)
}

func testGetContainerWithMostWater(t *testing.T) {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	assert.Equal(t, GetContainerWithMostWater(a), 49, "")
}

func testRomanToInteger(t *testing.T) {
	s := "XXVII"
	v := "LIX"
	assert.Equal(t, RomanToInteger(s), 27, "")
	assert.Equal(t, RomanToInteger(v), 59, "")
}

func testThreeSum(t *testing.T) {
	var nums = []int{-1, 0, 1, 2, -1, -4}
	results := ThreeSum(nums)
	assert.Equal(t, len(results), 2, "")
	assert.Equal(t, results[1], []int{-1, -1, 2}, "")
	assert.Equal(t, results[0], []int{-1, 0, 1}, "")
}
