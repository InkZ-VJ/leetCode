package linkedlist

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tempEven := head.Next
	even := head.Next
	ans := head
	for head != nil && head.Next != nil && tempEven != nil && tempEven.Next != nil {
		head.Next = head.Next.Next
		tempEven.Next = tempEven.Next.Next
		head = head.Next
		tempEven = tempEven.Next
	}
	head.Next = even
	return ans
}
