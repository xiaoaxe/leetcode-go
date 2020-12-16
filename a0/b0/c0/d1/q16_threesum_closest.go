// 找到距离target最近的三个值
// author: baoqiang
// time: 2020/12/15 10:11 下午
package d1

import (
	"math"
	"sort"

	"github.com/xiaoaxe/leetcode-go/common"
)

// 排序之后，前后夹逼
func ThreeSumClosest(nums []int, target int) int {
	if len(nums) <= 2 {
		// no need
		return math.MaxInt32
	}

	var result = math.MaxInt32
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		for j, k := i, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			diff := sum - target
			if diff == 0 {
				return sum
			} else if diff < 0 {
				j++
			} else {
				k--
			}

			if abs(diff) < result {
				result = sum
			}

		}
	}

	return result
}

func abs(a int) int {
	return common.IfElse(a > 0, a, -a)
}
