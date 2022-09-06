package solution

type Node struct {
	Val      int
	Children []*Node
}

func LevelOrder(root *Node) [][]int {
	var ans [][]int = make([][]int, 0)
	// dfsForLevelOrderOfNAry(root, 0, &ans)
	bfsForLevelOrderOfNAry(root, 0, ans)

	return ans

}

// 10ms
func dfsForLevelOrderOfNAry(root *Node, level int, ans *[][]int) {
	if root == nil {
		return
	}
	if len(*ans) <= level {
		(*ans) = append((*ans), []int{})
	}
	(*ans)[level] = append((*ans)[level], root.Val)
	for _, child := range root.Children {
		dfsForLevelOrderOfNAry(child, level+1, ans)
	}

}
func bfsForLevelOrderOfNAryWithNil(root *Node, level int, ans *[][]int) {
	var queue []*Node = make([]*Node, 0)
	queue = append(queue, root, nil)
	var temp []int = make([]int, 0)
	for len(queue) > 0 {
		pop := queue[0]
		queue = queue[1:]
		if pop == nil {
			if len(temp) > 0 {
				*ans = append(*ans, temp)
				temp = []int{}
				queue = append(queue, nil)
				continue
			}
		} else {
			temp = append(temp, pop.Val)
			for _, child := range pop.Children {
				queue = append(queue, child)
			}
		}

	}
}

func bfsForLevelOrderOfNAry(root *Node, level int, ans [][]int) {
	var queue []*Node = make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {

		n := len(queue)
		var temp []int = make([]int, 0)
		for i := 0; i < n; i++ {
			pop := queue[0]
			queue = queue[1:]
			temp = append(temp, pop.Val)
			if len(pop.Children) > 0 {
				queue = append(queue, pop.Children...)
			}
		}
		ans = append(ans, temp)
	}

}
