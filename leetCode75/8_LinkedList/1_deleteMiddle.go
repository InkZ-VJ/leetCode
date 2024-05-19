package linkedlist

func (head *ListNode) deleteMiddle() *ListNode {
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	if length == 1 {
		return nil
	}
	middle := length / 2
	node := head
	for i := 0; i < middle-1; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
	return head
}
