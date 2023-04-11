package solution

import (
	"github.com/imrajatmehta/leetcode/library"
)

type CloneGraphNode struct {
	Val       int
	Neighbors []*CloneGraphNode
}

func CloneGraph(node *CloneGraphNode) *CloneGraphNode {
	if node == nil {
		return node
	}
	hash := make(map[int]*CloneGraphNode)
	queue := library.Queue{}
	queue.Add(node)
	hash[node.Val] = &CloneGraphNode{Val: node.Val}
	for !queue.IsEmpty() {
		temp, _ := queue.Remove()
		pop := temp.(*CloneGraphNode)
		for i := range pop.Neighbors {
			child := pop.Neighbors[i]
			if _, ok := hash[child.Val]; !ok {
				queue.Add(child)
				hash[child.Val] = &CloneGraphNode{Val: child.Val}
			}
			hash[pop.Val].Neighbors = append(hash[pop.Val].Neighbors, hash[child.Val])
		}
	}
	return hash[1]
}
