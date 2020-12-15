// 最长不重复子串
// author: baoqiang
// time: 2020/12/14 9:01 下午
package d0

import "github.com/xiaoaxe/leetcode-go/common"

func GetLongestSubstringWithoutRepeating(s string) int {
	var result int
	var freq = make(map[uint8]int)
	var left, right int

	for left < len(s) {
		if right < len(s) && freq[s[right]] == 0 {
			freq[s[right]] = 1
			right++
		} else {
			freq[s[left]] = 0
			left++
		}
		result = max(result, right-left)
	}

	return result
}

func max(a, b int) int {
	return common.IfElse(a > b, a, b)
}
