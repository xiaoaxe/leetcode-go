// run test
// author: baoqiang
// time: 2020/12/14 7:51 下午
package d0

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestDemo(t *testing.T) {
	//testTwoSum(t)
	//testAddTwoNumbers2(t)
	//testGetLongestSubstringWithoutRepeating(t)
	//testReverseInteger(t)
	testIsPalindromeNumber(t)
}

func testTwoSum(t *testing.T) {
	array := []int64{3, 6, 1, 10}
	var sum int64 = 16
	idx0, idx1 := twoSum2(array, sum)
	assert.Equal(t, idx0, 1, "")
	assert.Equal(t, idx1, 3, "")
}

func testAddTwoNumbers(t *testing.T) {
	arr1 := []int{2, 4, 3}
	arr2 := []int{5, 6, 4}
	results := AddTwoNumbers(arr1, arr2)
	assert.Equal(t, len(results), 3, "")
	assert.Equal(t, results[0], 7, "")
	assert.Equal(t, results[1], 0, "")
	assert.Equal(t, results[2], 8, "")
}

func testAddTwoNumbers2(t *testing.T) {
	l1 := &LinkNode{2, &LinkNode{4, &LinkNode{5, nil}}}
	l2 := &LinkNode{5, &LinkNode{6, &LinkNode{4, nil}}}
	result := AddTwoNumbers2(l1, l2)

	var valStr string
	for result != nil {
		valStr += fmt.Sprintf("%d", result.val)
		result = result.next
	}

	assert.Equal(t, valStr, "7001", "")
}

func testGetLongestSubstringWithoutRepeating(t *testing.T) {
	//s := "pwwkew"
	s := "paakea"
	maxlen := GetLongestSubstringWithoutRepeating(s)
	assert.Equal(t, maxlen, 3, "")
}

func testReverseInteger(t *testing.T) {
	a := -12345
	b := ReverseInteger2(a)
	assert.Equal(t, b, -54321, "")
}

func testIsPalindromeNumber(t *testing.T) {
	a := 12345
	b := 12321
	assert.Equal(t, IsPalindromeNumber(a), false, "")
	assert.Equal(t, IsPalindromeNumber(b), true, "")
}
