// 数组围成的前后矩形，使得面积最大
// author: baoqiang
// time: 2020/12/15 8:52 下午
package d1

import "github.com/xiaoaxe/leetcode-go/common"

// 关键点: 只有"更矮的"移动才有可能使得容器面积变大
func GetContainerWithMostWater(a []int) int {
	var start, end = 0, len(a) - 1
	var result int

	for start < end {
		height := 0
		width := end - start
		if a[start] < a[end] {
			height = a[start]
			start++
		} else {
			height = a[end]
			end--
		}
		tmp := height * width
		result = common.IfElse(tmp > result, tmp, result)
	}

	return result
}
