// 两个排序数组的中位数
// author: baoqiang
// time: 2020/12/14 10:35 下午
package d0

import "github.com/xiaoaxe/leetcode-go/common"

// 二分查找
func GetMedianOfTwoSortedArrays(num1, num2 []int) int {
	var i, j int
	var mid = (len(num1) + len(num2)) / 2
	var hit1 bool

	for i < len(num1) || j < len(num2) {
		if num1[i] < num2[j] {
			hit1 = true
			i++
		} else {
			j++
		}

		if i+j == mid {
			return common.IfElse(hit1, num1[i], num2[j])
		}

	}

	return 0
}
