package solution

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	root := head
	ans := head
	var prev *ListNode
	for n > 0 {
		n--
		root = root.Next

	}
	for root != nil {

		root = root.Next
		prev = head
		head = head.Next
	}
	if prev == nil {
		ans = ans.Next
	} else {
		prev.Next = prev.Next.Next
	}

	return ans

}
