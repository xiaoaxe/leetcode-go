// 罗马数字转整数
// author: baoqiang
// time: 2020/12/15 9:28 下午
package d1

var romap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToInteger(s string) int {
	var result = 0
	var lastnum int

	for i := len(s) - 1; i >= 0; i-- {
		num := romap[s[i]]

		if num < lastnum {
			result -= num
		} else {
			result += num
		}

		lastnum = num
	}

	return result
}
