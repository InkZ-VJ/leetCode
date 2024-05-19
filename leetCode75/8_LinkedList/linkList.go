package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *ListNode) Length() int {
	length := 0
	for node := list; node != nil; node = node.Next {
		length++
	}
	return length
}
