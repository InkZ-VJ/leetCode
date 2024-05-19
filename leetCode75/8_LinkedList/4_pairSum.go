package linkedlist

func pairSum(head *ListNode) int {
	middle := Length(head) / 2

	currentList := head

	curr := head
	for i := 0; i < middle; i++ {
		curr = curr.Next
	}

	reverseList := reverseList(curr)

	maxSum := 0

	for i := 0; i < middle; i++ {
		sum := currentList.Val + reverseList.Val
		if sum > maxSum {
			maxSum = sum
		}
		currentList = currentList.Next
		reverseList = reverseList.Next
	}
	return maxSum
}

func Length(list *ListNode) int {
	length := 0
	for node := list; node != nil; node = node.Next {
		length++
	}
	return length
}
