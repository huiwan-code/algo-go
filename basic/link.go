package algo

type ListNode struct {
	next  		*ListNode
	val  		int
}
// 非递归反转链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		nextNode := head.next
		head.next = prev
		prev = head
		head = nextNode
	}
	return prev
}

// 递归反转链表
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	newHead := reverseList2(head.next)
	head.next.next = head
	head.next = nil
	return newHead
}
