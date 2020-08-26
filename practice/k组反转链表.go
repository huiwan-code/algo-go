package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	temp := head
	for i := 1; i < k && temp != nil; i ++ {
		temp = temp.Next
	}
	if temp == nil {
		return head
	}
	t2H := temp.Next
	temp.Next = nil
	newHead := reverse(head)
	head.Next = reverseKGroup(t2H, k)
	return newHead
}
func reverse(head *ListNode) *ListNode {
	if head.Next == nil || head == nil {
		return head
	}
	newHead := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
func main() {

}