// 列表，两个列表各个元素相加
// author: baoqiang
// time: 2020/12/14 8:11 下午
package d0

func AddTwoNumbers(nums1, nums2 []int) []int {
	var results []int

	size := len(nums1)
	carry := 0

	for i := 0; i < size; i++ {
		sum := nums1[i] + nums2[i] + carry
		if sum >= 10 {
			carry = 1
			results = append(results, sum-10)
		} else {
			carry = 0
			results = append(results, sum)
		}
	}

	return results
}

// 1. 定义链表节点
// 2. 处理长度不一致的情况
// 3. 使用头空节点
// 4. 处理边界情况，都为空了但是carry为1的情况
type LinkNode struct {
	val  int
	next *LinkNode
}

func AddTwoNumbers2(l1, l2 *LinkNode) *LinkNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	var head = &LinkNode{}
	var current = head
	var carry = 0
	for l1 != nil || l2 != nil {
		x, y := ifElse(l1 != nil, l1.val, 0), ifElse(l2 != nil, l2.val, 0)

		sum := x + y + carry
		val := ifElse(sum >= 10, sum-10, sum)
		carry = ifElse(sum >= 10, 1, 0)

		node := &LinkNode{
			val:  val,
			next: nil,
		}
		current.next = node
		current = node

		if l1 != nil {
			l1 = l1.next
		}
		if l2 != nil {
			l2 = l2.next
		}
	}

	//!!notice!!, carry may be 1
	if carry == 1 {
		node := &LinkNode{
			val:  1,
			next: nil,
		}
		current.next = node
		current = node
	}

	// !notice!!
	return head.next
}

func ifElse(cond bool, x, y int) int {
	if cond {
		return x
	}
	return y
}
