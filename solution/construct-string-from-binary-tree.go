package solution

import (
	"strconv"
	"strings"
)

func Tree2str(root *TreeNode) string {
	ans := []string{}
	dfsForTree2str(root, &ans)
	anss := strings.Join(ans, "")
	return anss[1 : len(anss)-1]
}
func dfsForTree2str(root *TreeNode, ans *[]string) {
	if root == nil {
		return
	}
	*ans = append(*ans, "(", strconv.Itoa(root.Val))

	dfsForTree2str(root.Left, ans)
	if root.Left == nil && root.Right != nil {
		*ans = append(*ans, "()")
	}
	dfsForTree2str(root.Right, ans)

	*ans = append(*ans, ")")
}
