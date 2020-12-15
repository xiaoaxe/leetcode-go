// 判断是否是回文数
// author: baoqiang
// time: 2020/12/15 8:42 下午
package d0

import "fmt"

func IsPalindromeNumber(a int) bool {
	if a < 0 {
		return false
	}
	return ReverseInteger2(a) == a
}

func IsPalindromeNumber2(a int) bool {
	if a < 0 {
		return false
	}

	s := fmt.Sprintf("%d", a)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
