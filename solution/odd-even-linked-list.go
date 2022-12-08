package solution

func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	oddNodes := head
	evenNodes := head.Next
	oddNodesLast := oddNodes
	evenNodesLast := evenNodes
	head = head.Next.Next

	even := true
	for head != nil {

		if even {
			oddNodesLast.Next = head
			oddNodesLast = head
		} else {
			evenNodesLast.Next = head
			evenNodesLast = head
		}

		even = !even
		head = head.Next
	}
	evenNodesLast.Next = nil
	oddNodesLast.Next = evenNodes
	return oddNodes
}
