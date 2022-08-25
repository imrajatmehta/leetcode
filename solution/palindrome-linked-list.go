package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

var left *ListNode

func IsPalindrome(head *ListNode) bool {
	left = head
	if ok := isPalindromeUtils(head); ok {
		return true
	}
	return false

}

func isPalindromeUtils(right *ListNode) bool {
	if right == nil {
		return true
	}
	flag := isPalindromeUtils(right.Next)
	if flag && left.Val == right.Val {
		left = left.Next
		return true
	}
	return false
}
