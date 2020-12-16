// 三个数和为0
// author: baoqiang
// time: 2020/12/15 9:42 下午
package d1

import "sort"

func ThreeSum(a []int) [][]int {
	var results [][]int

	var c = make(map[int]int)
	for _, v := range a {
		c[v] += 1
	}

	// sort a
	var keys []int
	for k := range c {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i, v1 := range keys {
		if v1 == 0 && c[v1] >= 3 {
			results = append(results, []int{v1, v1, v1})
		}

		for _, v2 := range keys[i+1:] {
			//if j <= i {
			//	continue
			//}

			if v1*2+v2 == 0 && c[v1] >= 2 && c[v2] >= 1 {
				results = append(results, []int{v1, v1, v2})
			}
			if v1+v2*2 == 0 && c[v1] >= 1 && c[v2] >= 2 {
				results = append(results, []int{v1, v2, v2})
			}

			// found one
			if c[-(v1+v2)] > 0 && -(v1+v2) > v2 {
				results = append(results, []int{v1, v2, -(v1 + v2)})
			}

		}
	}

	return results
}
