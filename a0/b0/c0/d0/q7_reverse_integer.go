// 反转整数
// author: baoqiang
// time: 2020/12/15 1:10 下午
package d0

import (
	"fmt"
	"strconv"
)

func ReverseInteger(a int) int {
	var t []byte

	s := fmt.Sprintf("%d", a)
	if s[0] == '-' {
		t = append(t, '-')
		s = s[1:]
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 0 {
			continue
		}
		t = append(t, s[i])
	}

	v, _ := strconv.ParseInt(string(t), 10, 64)
	return int(v)
}

// 不用转化成string，直接使用int反转
// 使用"1<<31-1"表示最大整数!!!
func ReverseInteger2(a int) int {
	var b int
	for a != 0 {
		b = b*10 + a%10
		a = a / 10
	}
	// [-2147483648,2147483647]
	if b > 1<<31-1 || b < -(1<<31) {
		return 0
	}

	return b
}
