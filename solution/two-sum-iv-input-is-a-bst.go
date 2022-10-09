package solution

import "github.com/imrajatmehta/leetcode/library/stack"

// Stack Based Solution
// T-O(N)
// S-O(N)
func FindTarget(root *TreeNode, k int) bool {

	leftStack := stack.New()
	rightStack := stack.New()

	iterator := root
	for iterator != nil {
		leftStack.Push(iterator)
		iterator = iterator.Left
	}

	iterator = root
	for iterator != nil {
		rightStack.Push(iterator)
		iterator = iterator.Right
	}

	for leftStack.Peek() != nil && rightStack.Peek() != nil && leftStack.Peek().(*TreeNode).Val != rightStack.Peek().(*TreeNode).Val {
		val1 := leftStack.Peek().(*TreeNode).Val
		val2 := rightStack.Peek().(*TreeNode).Val
		if val1+val2 == k {
			return true
		}

		if val1+val2 < k {
			pop := leftStack.Pop().(*TreeNode).Right
			for pop != nil {
				leftStack.Push(pop)
				pop = pop.Left
			}

		} else {
			pop := rightStack.Pop().(*TreeNode).Left
			for pop != nil {
				rightStack.Push(pop)
				pop = pop.Right
			}
		}

	}
	return false
}
