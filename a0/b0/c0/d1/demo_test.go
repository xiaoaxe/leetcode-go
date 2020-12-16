// test run
// author: baoqiang
// time: 2020/12/15 8:54 下午
package d1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDemo(t *testing.T) {
	//testGetContainerWithMostWater(t)
	//testRomanToInteger(t)
	//testThreeSum(t)
	//testGetLetterCombinations(t)
	testRemoveNthNodeFromEnd(t)
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

func testThreeSumClosest(t *testing.T) {
	var nums = []int{-1, 2, 1, -4}
	var target = 1
	assert.Equal(t, ThreeSumClosest(nums, target), 2, "")
}

func testGetLetterCombinations(t *testing.T) {
	//var num = "23"
	var num = "2312"
	results := GetLetterCombinations(num)
	assert.Equal(t, len(results), 27, "")
	fmt.Print(results)
}

func testRemoveNthNodeFromEnd(t *testing.T) {
	node := &LinkNode{1, &LinkNode{2, &LinkNode{3, &LinkNode{4, &LinkNode{}}}}}
	res := RemoveNthNodeFromEnd(node, 3)
	var vals []int
	assert.NotNil(t, res, "")
	for res.next != nil {
		vals = append(vals, res.val)
		res = res.next
	}
	assert.Equal(t, vals, []int{1, 2, 4}, "")
}
