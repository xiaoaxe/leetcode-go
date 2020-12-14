// 求出两个值加起来是给定值的索引
// author: baoqiang
// time: 2020/12/14 7:47 下午
package d0

func twoSum(arr []int64, sum int64) (idx0, idx1 int) {
	idxMap := make(map[int64]int)
	for idx, num := range arr {
		idxMap[num] = idx
	}

	for idx, num := range arr {
		if try, ok := idxMap[sum-num]; ok {
			return idx, try
		}
	}
	return -1, -1
}

// 1. 命名nums, target, m, another
// 2. 只需要遍历一次就可以
func twoSum2(nums []int64, target int64) (i, j int) {
	m := make(map[int64]int)
	for idx, num := range nums {
		another := target - num
		if try, ok := m[another]; ok {
			return try, idx
		}

		m[num] = idx
	}
	return -1, -1
}
