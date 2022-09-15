package solution

import "fmt"

func PseudoPalindromicPaths(root *TreeNode) int {

	freq := make([]int, 10)
	return dfsPseudoPalindromicPaths(root, freq)

}
func dfsPseudoPalindromicPaths(root *TreeNode, freq []int) int {
	res := 0
	freq[root.Val]++
	if root.Left == nil && root.Right == nil {
		if isPalindromeArrayWithFreq(freq) {
			res = 1
		}
	} else {
		if root.Left != nil {
			res += dfsPseudoPalindromicPaths(root.Left, freq)
		}
		if root.Right != nil {
			res += dfsPseudoPalindromicPaths(root.Right, freq)
		}

	}
	freq[root.Val]--
	return res
}
func isPalindromeArrayWithFreq(arr []int) bool {
	extraFreq := 0
	for _, val := range arr {
		if val%2 != 0 {
			extraFreq++
		}
	}
	fmt.Println(arr)
	return extraFreq <= 1
}
