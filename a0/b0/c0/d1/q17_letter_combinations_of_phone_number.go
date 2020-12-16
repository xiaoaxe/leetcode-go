// 给定数字，输出数字对应的所有字母的组合(电话上0-9)
// author: baoqiang
// time: 2020/12/16 7:57 下午
package d1

import "fmt"

var numMap = map[string]string{
	"0": " ",
	"1": "",
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
var res []string

func GetLetterCombinations(num string) []string {
	if len(num) == 0 {
		return res
	}
	dfs(num, 0, "")
	return res
}

// 设想只有两个字母的场景
func dfs(num string, index int, sub string) {
	if index == len(num) {
		res = append(res, sub)
		return
	}
	candi := numMap[fmt.Sprintf("%v", num[index:index+1])]
	if len(candi) == 0 {
		dfs(num, index+1, sub)
	}

	for i := 0; i < len(candi); i++ {
		dfs(num, index+1, sub+fmt.Sprintf("%v", candi[i:i+1]))
	}
}
