package solution

/*
a sorted list of start times could be used to check whether a time slot is available. but this would take nlogn time to perform each time a booking occurs. check if start/end time fits and then append to resort the list.

the other way is to use a LL that is able to insert in consant time, but it would take n to search where to place the event.

use a tree that sorts the start/end times
if the new event start <= end or new event end >= start that means there is an overlap to return false
if the new event start >= end (there is no overlap), so add event to the right if there is no right node.
if the new event end <= start (there is no overlap), so add event to the left
*/
type MyCalendar1 struct {
	Head *CalenderNode
}

func Constructor1() MyCalendar1 {
	return MyCalendar1{}
}

func (this *MyCalendar1) Book(start int, end int) bool {

	return checkAndInsertInBST(&this.Head, start, end)

}
func checkAndInsertInBST(head **CalenderNode, s, e int) bool {
	if *head == nil {
		*head = &CalenderNode{Start: s, End: e}
		return true
	} else if (*head).Start >= e {
		return checkAndInsertInBST(&(*head).Left, s, e)
	} else if (*head).End <= s {
		return checkAndInsertInBST(&(*head).Right, s, e)
	}
	return false
}

type CalenderNode struct {
	Start, End  int
	Left, Right *CalenderNode
}
