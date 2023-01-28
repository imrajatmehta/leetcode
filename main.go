package main

import (
	"fmt"

	"github.com/imrajatmehta/leetcode/solution"
)

func main() {
	// words := [][]int{{5, 1}, {5, 2}, {5, 3}, {6, 2}, {6, 3}, {7, 5}, {7, 6}}

	// a := []int{3, 5, 10, 7, 5, 3, 5, 5, 4, 8, 1}

	// five := solution.ListNode{Val: 5}
	// four := solution.ListNode{Val: 4}
	// three := solution.ListNode{Val: 3}
	// two := solution.ListNode{Val: 2}
	// one := solution.ListNode{Val: 1}
	// one.Next = &two
	// two.Next = &three
	// three.Next = &four
	// four.Next = &five
	o := solution.SummaryConstructor()
	o.AddNum(6)
	fmt.Println(o.GetIntervals())
	o.AddNum(6)
	fmt.Println(o.GetIntervals())
	o.AddNum(0)
	fmt.Println(o.GetIntervals())
	o.AddNum(4)
	fmt.Println(o.GetIntervals())
	o.AddNum(8)
	fmt.Println(o.GetIntervals())
	o.AddNum(7)
	fmt.Println(o.GetIntervals())
	o.AddNum(6)
	fmt.Println(o.GetIntervals())
	o.AddNum(4)
	fmt.Println(o.GetIntervals())
	o.AddNum(7)
	fmt.Println(o.GetIntervals())
	o.AddNum(5)
	fmt.Println(o.GetIntervals())
	o.AddNum(9)
	fmt.Println(o.GetIntervals())
	o.AddNum(1)
	fmt.Println(o.GetIntervals())

}
