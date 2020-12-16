// 移除结尾的第n个节点
// author: baoqiang
// time: 2020/12/16 8:39 下午
package d1

type LinkNode struct {
	val  int
	next *LinkNode
}

// !!!快慢指针!!!
func RemoveNthNodeFromEnd(head *LinkNode, n int) *LinkNode {
	if head == nil {
		return nil
	}
	var fast, slow = head, head

	for n > 0 {
		fast = fast.next
		n--
	}

	if fast == nil {
		return head.next
	}

	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}

	slow.next = slow.next.next
	return head
}
