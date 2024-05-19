package linkedlist

func reverseList(head *ListNode) *ListNode {
	var prev, current, nextNode *ListNode
	current = head

	for current != nil {
		nextNode = current.Next
		current.Next = prev
		prev = current
		current = nextNode
	}
	return prev
}
